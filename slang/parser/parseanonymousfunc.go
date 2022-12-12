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

	if p.curToken.Type() != tokens.TypeLPAREN {
		return nil, NewErrWrongTokenType(tokens.TypeLPAREN, p.curToken.Type())
	}

	p.nextToken()

	first := true
	argNames := []string{}

	for p.curToken.Type() != tokens.TypeRPAREN {
		if !first {
			if p.curToken.Type() != tokens.TypeCOMMA {
				return nil, NewErrWrongTokenType(tokens.TypeCOMMA, p.curToken.Type())
			}
			p.nextToken()
		}

		if p.curToken.Type() != tokens.TypeIDENT {
			return nil, NewErrWrongTokenType(tokens.TypeIDENT, p.curToken.Type())
		}

		argNames = append(argNames, p.curToken.Value())

		p.nextToken()

		first = false
	}

	p.nextToken()

	if p.curToken.Type() != tokens.TypeLBRACE {
		return nil, NewErrWrongTokenType(tokens.TypeLBRACE, p.curToken.Type())
	}

	p.nextToken()

	stmnts := []slang.Statement{}

	for p.curToken.Type() != tokens.TypeRBRACE {
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
