package parser

import (
	"github.com/jamestunnell/go-synth/slang"
	"github.com/jamestunnell/go-synth/slang/statements"
)

func (p *Parser) parseReturn() (slang.Statement, *ParseErr) {
	p.pushContext(slang.StatementRETURN.String())

	defer p.context.Pop()

	p.nextToken()

	expr, err := p.parseExpression()
	if err != nil {
		return nil, err
	}

	return statements.NewReturn(expr), nil
}
