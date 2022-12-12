package parser

import (
	"github.com/jamestunnell/go-synth/slang"
	"github.com/jamestunnell/go-synth/slang/statements"
	"github.com/jamestunnell/go-synth/slang/tokens"
)

func (p *Parser) parseAssign() (slang.Statement, error) {
	ident, err := p.parseIdentifier()
	if err != nil {
		return nil, err
	}

	p.nextToken()

	if err = p.curTokenMustBe(tokens.TypeASSIGN); err != nil {
		return nil, err
	}

	p.nextToken()

	value, err := p.parseExpression()
	if err != nil {
		return nil, err
	}

	return statements.NewAssign(ident, value), nil
}
