package parser

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/jamestunnell/go-synth/slang"
	"github.com/jamestunnell/go-synth/slang/expressions"
	"github.com/jamestunnell/go-synth/slang/statements"
	"github.com/jamestunnell/go-synth/slang/tokens"
)

type Parser struct {
	lexer slang.Lexer

	curToken  slang.Token
	peekToken slang.Token
}

var (
	errEmptyFuncBody = errors.New("function body is empty")
	errMissingReturn = errors.New("function missing return")
)

func New(l slang.Lexer) slang.Parser {
	p := &Parser{lexer: l}

	// Read two tokens, so curToken and peekToken are both set
	p.nextToken()

	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.lexer.NextToken()
}

func (p *Parser) ParseProgram() (*slang.Program, error) {
	program := slang.NewProgram()

	for p.curToken.Type() != tokens.TypeEOF {
		st, err := p.parseStatement()
		if err != nil {
			return nil, err
		}

		if st != nil {
			program.AddStatement(st)
		}

		p.nextToken()
	}

	return program, nil
}

func (p *Parser) parseStatement() (slang.Statement, error) {
	var s slang.Statement

	var err error

	switch p.curToken.Type() {
	case tokens.TypeLET:
		s, err = p.parseLet()
		if err != nil {
			return nil, fmt.Errorf("failed to parse let statement: %w", err)
		}
	case tokens.TypeRETURN:
		s, err = p.parseReturn()
		if err != nil {
			return nil, fmt.Errorf("failed to parse return statement: %w", err)
		}
	case tokens.TypeIF:
		s, err = p.parseIf()
		if err != nil {
			return nil, fmt.Errorf("failed to parse if statement: %w", err)
		}
	default:
		return nil, NewErrBadStatementStart(p.curToken)
	}

	return s, nil
}

func (p *Parser) parseLet() (slang.Statement, error) {
	p.nextToken()

	ident, err := p.parseIdentifier()
	if err != nil {
		return nil, err
	}

	p.nextToken()

	if p.curToken.Type() != tokens.TypeASSIGN {
		return nil, NewErrWrongTokenType(tokens.TypeASSIGN, p.curToken.Type())
	}

	p.nextToken()

	value, err := p.parseExpression()
	if err != nil {
		return nil, err
	}

	return statements.NewAssign(ident, value), nil
}

func (p *Parser) parseReturn() (slang.Statement, error) {
	p.nextToken()

	expr, err := p.parseExpression()
	if err != nil {
		err = fmt.Errorf("failed to parse return expression: %w", err)

		return nil, err
	}

	return statements.NewReturn(expr), nil
}

func (p *Parser) parseIf() (slang.Statement, error) {
	return nil, NewErrNotImplemented("if statement")
}

func (p *Parser) parseIdentifier() (*expressions.Identifier, error) {
	if p.curToken.Type() != tokens.TypeIDENT {
		return nil, NewErrWrongTokenType(tokens.TypeIDENT, p.curToken.Type())
	}

	return expressions.NewIdentifier(p.curToken.Value()), nil
}

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

		if op, success := AsBinaryOperator(p.peekToken); success {
			p.nextToken()

			p.nextToken()

			right, err := p.parseExpression()
			if err != nil {
				err = fmt.Errorf("failed to parse binary operation right-hand expression: %w", err)
				return nil, err
			}

			return expressions.NewBinaryOperation(op, val, right), nil
		}

		return val, nil

	case tokens.TypeLPAREN:
		return p.parseGroupedExpr()
	}

	return nil, NewErrBadExpressionStart(p.curToken)
}

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

func (p *Parser) parseGroupedExpr() (slang.Expression, error) {
	return nil, NewErrNotImplemented("grouped expression")
}

func (p *Parser) parseInteger(str string) (slang.Expression, error) {
	i, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("failed to parse '%s' as int: %w", str, err)
	}

	return expressions.NewInteger(i), nil
}

func (p *Parser) parseFloat(str string) (slang.Expression, error) {
	f, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return nil, fmt.Errorf("failed to parse '%s' as float: %w", str, err)
	}

	return expressions.NewFloat(f), nil
}
