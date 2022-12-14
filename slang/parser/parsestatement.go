package parser

import (
	"errors"

	"github.com/jamestunnell/go-synth/slang"
	"github.com/jamestunnell/go-synth/slang/statements"
)

var errBadStatementStart = errors.New("bad statment start")

func (p *Parser) parseStatement() (slang.Statement, *ParseErr) {
	var s slang.Statement

	var pErr *ParseErr

	switch p.curToken.Info.Type() {
	case slang.TokenRETURN:
		s, pErr = p.parseRetStatement()
	// case slang.TokenIF:
	// 	s, pErr = p.parseIfStatement()
	default:
		s, pErr = p.parseExprStatement()
	}

	// advance to semicolon or newline
	if p.peekTokenIs(slang.TokenSEMICOLON) || p.peekTokenIs(slang.TokenLINE) {
		p.nextToken()
	}

	return s, pErr
}

func (p *Parser) parseRetStatement() (slang.Statement, *ParseErr) {
	p.pushContext(slang.StatementRETURN.String())

	defer p.context.Pop()

	p.nextToken()

	expr := p.parseExpression(PrecedenceLOWEST)

	return statements.NewReturn(expr), nil
}

// func (p *Parser) parseIfStatement() slang.Statement {
// 	p.pushContext(slang.StatementIF.String())

// 	defer p.context.Pop()

// 	p.nextToken()

// 	cond := p.parseExpression(PrecedenceLOWEST)

// 	p.nextToken()

// 	if !curTokenMustBe(slang.TokenLBRACE); err != nil {
// 		return nil, err
// 	}

// 	p.nextTokenSkipLines()

// 	// this may generate errors that we don't see
// 	body := p.parseStatementsUntil(slang.TokenRBRACE)

// 	return statements.NewIf(cond, body...), nil
// }

func (p *Parser) parseExprStatement() (slang.Statement, *ParseErr) {
	p.pushContext(slang.StatementEXPRESSION.String())

	defer p.context.Pop()

	expr := p.parseExpression(PrecedenceLOWEST)

	return statements.NewExpression(expr), nil
}
