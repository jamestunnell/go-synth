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

	body := p.parseStatementsUntil(slang.TokenRBRACE)

	if len(body) == 0 {
		return nil, p.NewParseErr(errEmptyFuncBody)
	}

	last := body[len(body)-1]
	if last.Type() != slang.StatementRETURN {
		return nil, p.NewParseErr(errMissingReturn)
	}

	afunc := expressions.NewAnonymousFunction(argNames, body)

	return afunc, nil
}
