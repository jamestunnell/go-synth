package parser

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/jamestunnell/go-synth/slang"
	"github.com/jamestunnell/go-synth/slang/expressions"
)

var errBadExpressionStart = errors.New("bad expression start")

func (p *Parser) parseExpression(prec Precedence) slang.Expression {
	prefix := p.prefixParseFns[p.curToken.Info.Type()]
	if prefix == nil {
		err := NewErrMissingPrefixParseFn(p.curToken.Info.Type())

		p.Errors = append(p.Errors, p.NewParseErr(err))

		return nil
	}
	leftExp := prefix()

	return leftExp
}

func (p *Parser) parseIdentifier() slang.Expression {
	return expressions.NewIdentifier(p.curToken.Info.Value())
}

func (p *Parser) parseTrue() slang.Expression {
	return expressions.NewBool(true)
}

func (p *Parser) parseFalse() slang.Expression {
	return expressions.NewBool(false)
}

func (p *Parser) parseNegative() slang.Expression {
	p.nextToken()

	val := p.parseExpression(PrecedencePREFIX)

	return expressions.NewNegative(val)
}

func (p *Parser) parseNot() slang.Expression {
	p.nextToken()

	val := p.parseExpression(PrecedencePREFIX)

	return expressions.NewNot(val)
}

func (p *Parser) parseInteger() slang.Expression {
	str := p.curToken.Info.Value()

	i, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		err = fmt.Errorf("failed to parse '%s' as int: %w", str, err)

		p.Errors = append(p.Errors, p.NewParseErr(err))

		return nil
	}

	return expressions.NewInteger(i)
}

func (p *Parser) parseFloat() slang.Expression {
	str := p.curToken.Info.Value()

	f, err := strconv.ParseFloat(str, 64)
	if err != nil {
		err = fmt.Errorf("failed to parse '%s' as float: %w", str, err)

		p.Errors = append(p.Errors, p.NewParseErr(err))

		return nil
	}

	return expressions.NewFloat(f)
}

// 	switch p.curToken.Info.Type() {
// 	case slang.TokenFUNC:
// 		return p.parseAnonymousFunc()
// 	case slang.TokenIDENT, slang.TokenINT, slang.TokenFLOAT:
// 		var val slang.Expression
// 		var err error

// 		switch p.curToken.Info.Type() {
// 		case slang.TokenIDENT:
// 			val = expressions.NewIdentifier(p.curToken.Info.Value())
// 		case slang.TokenINT:
// 			val, err = p.parseInteger(p.curToken.Info.Value())
// 		case slang.TokenFLOAT:
// 			val, err = p.parseFloat(p.curToken.Info.Value())
// 		}

// 		if err != nil {
// 			return nil, NewParseError(err, p.curToken, p.currentContext())
// 		}

// 		if operator, opPrec, success := AsBinaryOperator(p.peekToken.Info.Type()); success {
// 			p.nextToken()

// 			p.nextToken()

// 			right, pErr := p.parseExpression(opPrec)
// 			if err != nil {
// 				return nil, pErr
// 			}

// 			return operator.MakeExpression(val, right), nil
// 		}

// 		return val, nil

// 		// case slang.TokenLPAREN:
// 		// 	return p.parseGroupedExpr()
// 	}

// 	err := errBadExpressionStart

// 	return nil, NewParseError(err, p.curToken, p.currentContext())
// }
