package lexer

import (
	"io"
	"unicode"

	"github.com/rs/zerolog/log"

	"github.com/jamestunnell/go-synth/slang"
	"github.com/jamestunnell/go-synth/slang/tokens"
)

type Lexer struct {
	scanner io.RuneScanner
	ch      rune
	str     string
}

func New(scanner io.RuneScanner) slang.Lexer {
	return &Lexer{
		scanner: scanner,
		ch:      0,
		str:     "",
	}
}

func (l *Lexer) NextToken() slang.Token {
	var tok slang.Token

	l.readRune()

	for isSpaceOrTab(l.ch) {
		l.readRune()
	}

	switch l.ch {
	case '\n', '\v', '\f':
		tok = tokens.LINE()
	case '>', '<', '=', '.', ',', ';', '(', ')', '{', '}', '+', '-', '*', '/':
		tok = l.readSymbol()
	case 0:
		tok = tokens.EOF()
	default:
		if isLetterOrUnderscore(l.ch) {
			tok = l.readIdentOrKeyword()
		} else if unicode.IsDigit(l.ch) {
			tok = l.readNumber()
		} else {
			tok = tokens.ILLEGAL(l.ch)
		}
	}

	return tok
}

func (l *Lexer) readRune() {
	r, _, _ := l.scanner.ReadRune()

	l.ch = r
	l.str = string([]rune{r})
}

func (l *Lexer) unreadRune() {
	if l.ch == 0 {
		return
	}

	err := l.scanner.UnreadRune()
	if err != nil {
		log.Warn().Err(err).Msg("Lexer: failed to unread rune")
	}
}

func (l *Lexer) readSymbol() slang.Token {
	var tok slang.Token

	switch l.ch {
	case '!':
		tok = l.readNot()
	case '<':
		tok = l.readLess()
	case '>':
		tok = l.readGreater()
	case '=':
		tok = l.readEqual()
	case '+':
		tok = l.readPlus()
	case '-':
		tok = l.readMinus()
	case '*':
		tok = l.readStar()
	case '/':
		tok = l.readSlash()
	case '.':
		tok = tokens.DOT()
	case ',':
		tok = tokens.COMMA()
	case ';':
		tok = tokens.SEMICOLON()
	case '(':
		tok = tokens.LPAREN()
	case ')':
		tok = tokens.RPAREN()
	case '{':
		tok = tokens.LBRACE()
	case '}':
		tok = tokens.RBRACE()
	default:
		log.Fatal().
			Str("rune", string([]rune{l.ch})).
			Msg("unexpected symbol rune")
	}

	return tok
}

func (l *Lexer) readNot() slang.Token {
	l.readRune()
	if l.ch == '=' {
		return tokens.NOTEQUAL()
	}

	l.unreadRune()
	return tokens.NOT()
}

func (l *Lexer) readLess() slang.Token {
	l.readRune()
	if l.ch == '=' {
		return tokens.LESSEQUAL()
	}

	l.unreadRune()
	return tokens.LESS()
}

func (l *Lexer) readGreater() slang.Token {
	l.readRune()
	if l.ch == '=' {
		return tokens.GREATEREQUAL()
	}

	l.unreadRune()
	return tokens.GREATER()
}

func (l *Lexer) readEqual() slang.Token {
	l.readRune()
	if l.ch == '=' {
		return tokens.EQUAL()
	}

	l.unreadRune()
	return tokens.ASSIGN()
}

func (l *Lexer) readPlus() slang.Token {
	l.readRune()
	switch l.ch {
	case '=':
		return tokens.PLUSEQUAL()
	case '+':
		return tokens.PLUSPLUS()
	}

	l.unreadRune()
	return tokens.PLUS()
}

func (l *Lexer) readMinus() slang.Token {
	l.readRune()
	switch l.ch {
	case '=':
		return tokens.MINUSEQUAL()
	case '-':
		return tokens.MINUSMINUS()
	}

	l.unreadRune()
	return tokens.MINUS()
}

func (l *Lexer) readStar() slang.Token {
	l.readRune()
	if l.ch == '=' {
		return tokens.STAREQUAL()
	}

	l.unreadRune()
	return tokens.STAR()
}

func (l *Lexer) readSlash() slang.Token {
	l.readRune()
	if l.ch == '=' {
		return tokens.SLASHEQUAL()
	}

	l.unreadRune()
	return tokens.SLASH()
}

func (l *Lexer) readIdentOrKeyword() slang.Token {
	runes := []rune{l.ch}

	l.readRune()

	for isLetterOrUnderscore(l.ch) {
		runes = append(runes, l.ch)

		l.readRune()
	}

	l.unreadRune()

	str := string(runes)

	switch str {
	case tokens.StrELSE:
		return tokens.ELSE()
	case tokens.StrFALSE:
		return tokens.FALSE()
	case tokens.StrFUNC:
		return tokens.FUNC()
	case tokens.StrIF:
		return tokens.IF()
	case tokens.StrRETURN:
		return tokens.RETURN()
	case tokens.StrTRUE:
		return tokens.TRUE()
	}

	return tokens.IDENT(str)
}

func (l *Lexer) readNumber() slang.Token {
	runes := []rune{l.ch}
	dotUsed := false

	l.readRune()

	for unicode.IsDigit(l.ch) || (l.ch == '.' && !dotUsed) {
		runes = append(runes, l.ch)

		if l.ch == '.' {
			dotUsed = true
		}

		l.readRune()
	}

	l.unreadRune()

	str := string(runes)

	if dotUsed {
		return tokens.FLOAT(str)
	}

	return tokens.INT(str)
}

func isLetterOrUnderscore(r rune) bool {
	return unicode.IsLetter(r) || r == '_'
}

func isSpaceOrTab(r rune) bool {
	return r == ' ' || r == '\t' || r == '\r'
}
