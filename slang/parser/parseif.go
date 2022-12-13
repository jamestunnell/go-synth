package parser

import (
	"github.com/jamestunnell/go-synth/slang"
	"github.com/jamestunnell/go-synth/slang/statements"
)

func (p *Parser) parseIf() (slang.Statement, *ParseErr) {
	p.pushContext(slang.StatementIF.String())

	defer p.context.Pop()

	p.nextToken()

	cond, err := p.parseExpression()
	if err != nil {
		return nil, err
	}

	p.nextToken()

	if err = p.curTokenMustBe(slang.TokenLBRACE); err != nil {
		return nil, err
	}

	p.nextTokenSkipLines()

	// this may generate errors that we don't see
	body := p.parseStatementsUntil(slang.TokenRBRACE)

	return statements.NewIf(cond, body...), nil
}
