package parser

import (
	"fmt"

	"github.com/jamestunnell/go-synth/slang"
	"github.com/jamestunnell/go-synth/slang/statements"
)

func (p *Parser) parseIf() (slang.Statement, error) {
	p.nextToken()

	cond, err := p.parseExpression()
	if err != nil {
		err = fmt.Errorf("failed to parse if condition expression: %w", err)

		return nil, err
	}

	p.nextToken()

	if err := p.curTokenMustBe(slang.TokenLBRACE); err != nil {
		return nil, err
	}

	p.nextTokenSkipLines()

	body, err := p.parseStatementsUntil(slang.TokenRBRACE)
	if err != nil {
		return nil, fmt.Errorf("failed to parse if body: %w", err)
	}

	return statements.NewIf(cond, body...), nil
}
