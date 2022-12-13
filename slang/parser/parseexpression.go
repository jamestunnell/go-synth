package parser

import (
	"errors"

	"github.com/jamestunnell/go-synth/slang"
	"github.com/jamestunnell/go-synth/slang/expressions"
)

var errBadExpressionStart = errors.New("bad expression start")

func (p *Parser) parseExpression() (slang.Expression, *ParseErr) {
	switch p.curToken.Info.Type() {
	case slang.TokenFUNC:
		return p.parseAnonymousFunc()
	case slang.TokenIDENT, slang.TokenINT, slang.TokenFLOAT:
		var val slang.Expression
		var err error

		switch p.curToken.Info.Type() {
		case slang.TokenIDENT:
			val = expressions.NewIdentifier(p.curToken.Info.Value())
		case slang.TokenINT:
			val, err = p.parseInteger(p.curToken.Info.Value())
		case slang.TokenFLOAT:
			val, err = p.parseFloat(p.curToken.Info.Value())
		}

		if err != nil {
			return nil, NewParseError(err, p.curToken, p.currentContext())
		}

		if operator, success := AsBinaryOperator(p.peekToken.Info.Type()); success {
			p.nextToken()

			p.nextToken()

			right, pErr := p.parseExpression()
			if err != nil {
				return nil, pErr
			}

			return operator.MakeExpression(val, right), nil
		}

		return val, nil

	case slang.TokenLPAREN:
		return p.parseGroupedExpr()
	}

	err := errBadExpressionStart

	return nil, NewParseError(err, p.curToken, p.currentContext())
}
