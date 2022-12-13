package lexer

import (
	"strings"

	"github.com/jamestunnell/go-synth/slang"
)

func ScanString(input string) []*slang.Token {
	l := New(strings.NewReader(input))

	toks := []*slang.Token{}

	for tok := l.NextToken(); tok.Info.Type() != slang.TokenEOF; tok = l.NextToken() {
		toks = append(toks, tok)
	}

	return toks
}
