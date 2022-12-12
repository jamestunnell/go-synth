package parser

import (
	"fmt"

	"github.com/jamestunnell/go-synth/slang"
	"github.com/jamestunnell/go-synth/slang/expressions"
)

func (p *Parser) parseExpression() (slang.Expression, error) {
	switch p.curToken.Type() {
	case slang.TokenFUNC:
		return p.parseAnonymousFunc()
	case slang.TokenIDENT, slang.TokenINT, slang.TokenFLOAT:
		var val slang.Expression
		var err error

		switch p.curToken.Type() {
		case slang.TokenIDENT:
			val = expressions.NewIdentifier(p.curToken.Value())
		case slang.TokenINT:
			val, err = p.parseInteger(p.curToken.Value())
		case slang.TokenFLOAT:
			val, err = p.parseFloat(p.curToken.Value())
		}

		if err != nil {
			return nil, fmt.Errorf("failed to parse expression: %w", err)
		}

		if operator, success := AsBinaryOperator(p.peekToken); success {
			p.nextToken()

			p.nextToken()

			right, err := p.parseExpression()
			if err != nil {
				err = fmt.Errorf("failed to parse binary operation right-hand expression: %w", err)
				return nil, err
			}

			return operator.MakeExpression(val, right), nil
		}

		return val, nil

	case slang.TokenLPAREN:
		return p.parseGroupedExpr()
	}

	return nil, NewErrBadExpressionStart(p.curToken)
}
