package parser

import (
	"github.com/jamestunnell/go-synth/slang"
	"github.com/jamestunnell/go-synth/slang/expressions"
)

func (p *Parser) parseAnonymousFunc() (slang.Expression, *ParseErr) {
	p.pushContext(slang.ExprANONYMOUSFUNC.String())

	defer p.context.Pop()

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

		argNames = append(argNames, p.curToken.Info.Value())

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
			return nil, err
		}

		stmnts = append(stmnts, s)
	}

	if len(stmnts) == 0 {
		return nil, p.NewParseErr(errEmptyFuncBody)
	}

	last := stmnts[len(stmnts)-1]
	if last.Type() != slang.StatementRETURN {
		return nil, p.NewParseErr(errMissingReturn)
	}

	afunc := expressions.NewAnonymousFunction(argNames, stmnts)

	return afunc, nil
}
