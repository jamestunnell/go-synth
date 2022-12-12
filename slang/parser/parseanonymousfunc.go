package parser

import (
	"fmt"

	"github.com/jamestunnell/go-synth/slang"
	"github.com/jamestunnell/go-synth/slang/expressions"
	"github.com/jamestunnell/go-synth/slang/statements"
)

func (p *Parser) parseAnonymousFunc() (slang.Expression, error) {
	p.nextToken()

	if err := p.curTokenMustBe(slang.TokenLPAREN); err != nil {
		return nil, err
	}

	p.nextToken()

	first := true
	argNames := []string{}

	for !p.curTokenIs(slang.TokenRPAREN) {
		if !first {
			if err := p.curTokenMustBe(slang.TokenCOMMA); err != nil {
				return nil, err
			}
			p.nextToken()
		}

		if err := p.curTokenMustBe(slang.TokenIDENT); err != nil {
			return nil, err
		}

		argNames = append(argNames, p.curToken.Value())

		p.nextToken()

		first = false
	}

	p.nextToken()

	if err := p.curTokenMustBe(slang.TokenLBRACE); err != nil {
		return nil, err
	}

	p.nextToken()

	stmnts := []slang.Statement{}

	for !p.curTokenIs(slang.TokenRBRACE) {
		s, err := p.parseStatement()
		if err != nil {
			return nil, fmt.Errorf("failed to parse func statement: %w", err)
		}

		stmnts = append(stmnts, s)
	}

	if len(stmnts) == 0 {
		return nil, errEmptyFuncBody
	}

	last := stmnts[len(stmnts)-1]
	if last.Type() != statements.TypeRETURN {
		return nil, errMissingReturn
	}

	afunc := expressions.NewAnonymousFunction(argNames, stmnts)

	return afunc, nil
}
