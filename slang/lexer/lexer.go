package lexer

import (
	"io"
	"unicode"

	"github.com/rs/zerolog/log"

	"github.com/jamestunnell/go-synth/slang"
	"github.com/jamestunnell/go-synth/slang/tokens"
)

type Lexer interface {
	NextToken() slang.Token
}

type lexer struct {
	scanner io.RuneScanner
	ch      rune
	str     string
}

func New(scanner io.RuneScanner) Lexer {
	l := &lexer{
		scanner: scanner,
	}

	l.readRune()

	return l
}

func (l *lexer) NextToken() slang.Token {
	var tok slang.Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		tok = tokens.ASSIGN()

		l.readRune()
	case ';':
		tok = tokens.SEMICOLON()

		l.readRune()
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

func (l *lexer) skipWhitespace() {
	for unicode.IsSpace(l.ch) {
		l.readRune()
	}
}

func (l *lexer) readRune() {
	r, _, _ := l.scanner.ReadRune()

	l.ch = r
	l.str = string([]rune{r})
}

func (l *lexer) unreadRune() {
	err := l.scanner.UnreadRune()
	if err != nil {
		log.Warn().Err(err).Msg("lexer: failed to unread rune")
	}
}

func (l *lexer) readIdentOrKeyword() slang.Token {
	runes := []rune{l.ch}

	l.readRune()

	for isLetterOrUnderscore(l.ch) {
		runes = append(runes, l.ch)

		l.readRune()
	}

	str := string(runes)

	switch str {
	case tokens.StrFUNC:
		return tokens.FUNC()
	case tokens.StrLET:
		return tokens.LET()
	}

	return tokens.IDENT(str)
}

func (l *lexer) readNumber() slang.Token {
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

	str := string(runes)

	if dotUsed {
		return tokens.FLOAT(str)
	}

	return tokens.INT(str)
}

func isLetterOrUnderscore(r rune) bool {
	return unicode.IsLetter(r) || r == '_'
}
