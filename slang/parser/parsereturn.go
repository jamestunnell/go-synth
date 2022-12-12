package parser

import (
	"fmt"

	"github.com/jamestunnell/go-synth/slang"
	"github.com/jamestunnell/go-synth/slang/statements"
)

func (p *Parser) parseReturn() (slang.Statement, ParseErr) {
	p.nextToken()

	expr, err := p.parseExpression()
	if err != nil {
		err = fmt.Errorf("failed to parse return expression: %w", err)

		return nil, err
	}

	return statements.NewReturn(expr), nil
}
