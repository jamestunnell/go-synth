package lexer

import (
	"strings"

	"github.com/jamestunnell/go-synth/slang"
	"github.com/jamestunnell/go-synth/slang/tokens"
)

func ScanString(input string) []slang.Token {
	l := New(strings.NewReader(input))

	toks := []slang.Token{}

	for tok := l.NextToken(); tok.Type() != tokens.TypeEOF; tok = l.NextToken() {
		toks = append(toks, tok)
	}

	return toks
}
