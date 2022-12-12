package parser

import (
	"github.com/jamestunnell/go-synth/slang/expressions"
	"github.com/jamestunnell/go-synth/slang/tokens"
)

func (p *Parser) parseIdentifier() (*expressions.Identifier, error) {
	if err := p.curTokenMustBe(tokens.TypeIDENT); err != nil {
		return nil, err
	}

	return expressions.NewIdentifier(p.curToken.Value()), nil
}
