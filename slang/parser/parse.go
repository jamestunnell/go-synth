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

	if err := p.Run(); err != nil {
		return nil, err
	}

	return slang.NewProgram(p.Statements...), nil
}
