package parser

import (
	"github.com/jamestunnell/go-synth/slang"
	"github.com/jamestunnell/go-synth/slang/statements"
)

func (p *Parser) parseAssign() (slang.Statement, *ParseErr) {
	p.pushContext(slang.StatementASSIGN.String())

	defer p.context.Pop()

	ident, err := p.parseIdentifier()
	if err != nil {
		return nil, err
	}

	p.nextToken()

	if err = p.curTokenMustBe(slang.TokenASSIGN); err != nil {
		return nil, err
	}

	p.nextToken()

	value, err := p.parseExpression()
	if err != nil {
		return nil, err
	}

	return statements.NewAssign(ident, value), nil
}
