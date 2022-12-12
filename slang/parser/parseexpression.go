package parser

import (
	"fmt"

	"github.com/jamestunnell/go-synth/slang"
	"github.com/jamestunnell/go-synth/slang/expressions"
	"github.com/jamestunnell/go-synth/slang/tokens"
)

func (p *Parser) parseExpression() (slang.Expression, error) {
	switch p.curToken.Type() {
	case tokens.TypeFUNC:
		return p.parseAnonymousFunc()
	case tokens.TypeIDENT, tokens.TypeINT, tokens.TypeFLOAT:
		var val slang.Expression
		var err error

		switch p.curToken.Type() {
		case tokens.TypeIDENT:
			val = expressions.NewIdentifier(p.curToken.Value())
		case tokens.TypeINT:
			val, err = p.parseInteger(p.curToken.Value())
		case tokens.TypeFLOAT:
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

	case tokens.TypeLPAREN:
		return p.parseGroupedExpr()
	}

	return nil, NewErrBadExpressionStart(p.curToken)
}
