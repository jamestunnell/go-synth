package parser

import (
	"fmt"

	"github.com/jamestunnell/go-synth/slang"
	"github.com/jamestunnell/go-synth/slang/expressions"
	"github.com/jamestunnell/go-synth/slang/statements"
	"github.com/jamestunnell/go-synth/slang/tokens"
)

func (p *Parser) parseAnonymousFunc() (slang.Expression, error) {
	p.nextToken()

	if err := p.curTokenMustBe(tokens.TypeLPAREN); err != nil {
		return nil, err
	}

	p.nextToken()

	first := true
	argNames := []string{}

	for !p.curTokenIs(tokens.TypeRPAREN) {
		if !first {
			if err := p.curTokenMustBe(tokens.TypeCOMMA); err != nil {
				return nil, err
			}
			p.nextToken()
		}

		if err := p.curTokenMustBe(tokens.TypeIDENT); err != nil {
			return nil, err
		}

		argNames = append(argNames, p.curToken.Value())

		p.nextToken()

		first = false
	}

	p.nextToken()

	if err := p.curTokenMustBe(tokens.TypeLBRACE); err != nil {
		return nil, err
	}

	p.nextToken()

	stmnts := []slang.Statement{}

	for !p.curTokenIs(tokens.TypeRBRACE) {
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
