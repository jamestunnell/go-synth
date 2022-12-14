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

	for !p.peekTokenIs(slang.TokenSEMICOLON) && prec < p.peekPrecedence() {
		infix := p.infixParseFns[p.peekToken.Info.Type()]
		if infix == nil {
			return leftExp
		}

		p.nextToken()

		leftExp = infix(leftExp)
	}

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
	return p.parsePrefixExpr(expressions.NewNegative)
}

func (p *Parser) parseNot() slang.Expression {
	return p.parsePrefixExpr(expressions.NewNot)
}

type newPrefixExprFn func(slang.Expression) slang.Expression

func (p *Parser) parsePrefixExpr(fn newPrefixExprFn) slang.Expression {
	p.nextToken()

	val := p.parseExpression(PrecedencePREFIX)

	return fn(val)
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

func (p *Parser) parseAdd(left slang.Expression) slang.Expression {
	return p.parseInfixExpr(left, expressions.NewAdd)
}

func (p *Parser) parseSubtract(left slang.Expression) slang.Expression {
	return p.parseInfixExpr(left, expressions.NewSubtract)
}

func (p *Parser) parseMultiply(left slang.Expression) slang.Expression {
	return p.parseInfixExpr(left, expressions.NewMultiply)
}

func (p *Parser) parseDivide(left slang.Expression) slang.Expression {
	return p.parseInfixExpr(left, expressions.NewDivide)
}

func (p *Parser) parseEqual(left slang.Expression) slang.Expression {
	return p.parseInfixExpr(left, expressions.NewEqual)
}

func (p *Parser) parseNotEqual(left slang.Expression) slang.Expression {
	return p.parseInfixExpr(left, expressions.NewNotEqual)
}

func (p *Parser) parseLess(left slang.Expression) slang.Expression {
	return p.parseInfixExpr(left, expressions.NewLess)
}

func (p *Parser) parseGreater(left slang.Expression) slang.Expression {
	return p.parseInfixExpr(left, expressions.NewGreater)
}

type newInfixExprFn func(left, right slang.Expression) slang.Expression

func (p *Parser) parseInfixExpr(left slang.Expression, fn newInfixExprFn) slang.Expression {
	prec := p.curPrecedence()

	p.nextToken()

	right := p.parseExpression(prec)

	return fn(left, right)
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
