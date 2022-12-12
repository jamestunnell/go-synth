package parser

import (
	"github.com/jamestunnell/go-synth/slang"
	"github.com/jamestunnell/go-synth/slang/expressions"
)

func (p *Parser) parseIdentifier() (*expressions.Identifier, error) {
	if err := p.curTokenMustBe(slang.TokenIDENT); err != nil {
		return nil, err
	}

	return expressions.NewIdentifier(p.curToken.Value()), nil
}
