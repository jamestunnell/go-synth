package parser

import (
	"fmt"

	"github.com/jamestunnell/go-synth/slang"
	"github.com/jamestunnell/go-synth/slang/statements"
	"github.com/jamestunnell/go-synth/slang/tokens"
)

func (p *Parser) parseIf() (slang.Statement, error) {
	p.nextToken()

	cond, err := p.parseExpression()
	if err != nil {
		err = fmt.Errorf("failed to parse if condition expression: %w", err)

		return nil, err
	}

	p.nextToken()

	if err := p.checkCurToken(tokens.TypeLBRACE); err != nil {
		return nil, err
	}

	p.nextTokenSkipLines()

	body, err := p.parseStatementsUntil(tokens.TypeRBRACE)
	if err != nil {
		return nil, fmt.Errorf("failed to parse if body: %w", err)
	}

	return statements.NewIf(cond, body...), nil
}
