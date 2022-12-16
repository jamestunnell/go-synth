package lexer

import (
	"io"
	"unicode"

	"github.com/rs/zerolog/log"

	"github.com/jamestunnell/go-synth/slang"
	"github.com/jamestunnell/go-synth/slang/tokens"
)

type Lexer struct {
	scanner   io.RuneScanner
	ch        rune
	str       string
	line, col int
}

func New(scanner io.RuneScanner) slang.Lexer {
	return &Lexer{
		scanner: scanner,
		ch:      0,
		str:     "",
		line:    1,
		col:     0,
	}
}

func (l *Lexer) NextToken() *slang.Token {
	var tokInfo slang.TokenInfo

	l.readRune()

	for isSpaceOrTab(l.ch) {
		// if isNewlineish(l.ch) {
		// 	l.line++
		// }

		l.readRune()
	}

	loc := slang.SourceLocation{
		Line:   l.line,
		Column: l.col,
	}

	switch l.ch {
	case '\n', '\v', '\f':
		tokInfo = tokens.LINE()

		l.line++
		l.col = 0
	case '!', '>', '<', '=', '.', ',', ';', '(', ')', '{', '}', '+', '-', '*', '/':
		tokInfo = l.readSymbol()
	case 0:
		tokInfo = tokens.EOF()
	default:
		if isLetterOrUnderscore(l.ch) {
			tokInfo = l.readIdentOrKeyword()
		} else if unicode.IsDigit(l.ch) {
			tokInfo = l.readNumber()
		} else {
			tokInfo = tokens.ILLEGAL(l.ch)
		}
	}

	return slang.NewToken(tokInfo, loc)
}

func (l *Lexer) readRune() {
	r, _, _ := l.scanner.ReadRune()

	l.col++
	l.ch = r
	l.str = string([]rune{r})
}

func (l *Lexer) unreadRune() {
	// nothing read yet
	if l.ch == 0 {
		return
	}

	err := l.scanner.UnreadRune()
	if err != nil {
		log.Warn().Err(err).Msg("Lexer: failed to unread rune")
	}

	l.col--
}

func (l *Lexer) readSymbol() slang.TokenInfo {
	var tok slang.TokenInfo

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

func (l *Lexer) readNot() slang.TokenInfo {
	l.readRune()
	if l.ch == '=' {
		return tokens.NOTEQUAL()
	}

	l.unreadRune()
	return tokens.BANG()
}

func (l *Lexer) readLess() slang.TokenInfo {
	l.readRune()
	if l.ch == '=' {
		return tokens.LESSEQUAL()
	}

	l.unreadRune()
	return tokens.LESS()
}

func (l *Lexer) readGreater() slang.TokenInfo {
	l.readRune()
	if l.ch == '=' {
		return tokens.GREATEREQUAL()
	}

	l.unreadRune()
	return tokens.GREATER()
}

func (l *Lexer) readEqual() slang.TokenInfo {
	l.readRune()
	if l.ch == '=' {
		return tokens.EQUAL()
	}

	l.unreadRune()
	return tokens.ASSIGN()
}

func (l *Lexer) readPlus() slang.TokenInfo {
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

func (l *Lexer) readMinus() slang.TokenInfo {
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

func (l *Lexer) readStar() slang.TokenInfo {
	l.readRune()
	if l.ch == '=' {
		return tokens.STAREQUAL()
	}

	l.unreadRune()
	return tokens.STAR()
}

func (l *Lexer) readSlash() slang.TokenInfo {
	l.readRune()
	if l.ch == '=' {
		return tokens.SLASHEQUAL()
	}

	l.unreadRune()
	return tokens.SLASH()
}

func (l *Lexer) readIdentOrKeyword() slang.TokenInfo {
	runes := []rune{l.ch}

	l.readRune()

	for unicode.IsDigit(l.ch) || isLetterOrUnderscore(l.ch) {
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

func (l *Lexer) readNumber() slang.TokenInfo {
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

func isNewlineish(r rune) bool {
	return r == '\n' || r == '\v' || r == '\f'
}
