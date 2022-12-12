package parser

import (
	"strings"

	"github.com/jamestunnell/go-synth/slang"
	"github.com/jamestunnell/go-synth/slang/lexer"
)

func Parse(input string) (*slang.Program, error) {
	r := strings.NewReader(input)
	l := lexer.New(r)
	p := New(l)

	return p.ParseProgram()
}
