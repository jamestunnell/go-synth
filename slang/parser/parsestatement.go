package parser

import (
	"errors"

	"github.com/jamestunnell/go-synth/slang"
	"github.com/jamestunnell/go-synth/slang/expressions"
	"github.com/jamestunnell/go-synth/slang/statements"
)

var errBadStatementStart = errors.New("bad statment start")

func (p *Parser) parseStatement() slang.Statement {
	var s slang.Statement

	switch p.curToken.Info.Type() {
	case slang.TokenRETURN:
		s = p.parseRetStatement()
	case slang.TokenIDENT:
		if p.peekTokenIs(slang.TokenASSIGN) {
			s = p.parseAssignStatement()
		} else {
			s = p.parseExprStatement()
		}
	default:
		s = p.parseExprStatement()
	}

	// advance to semicolon or newline
	if p.peekTokenIs(slang.TokenSEMICOLON) || p.peekTokenIs(slang.TokenLINE) {
		p.nextToken()
	}

	return s
}

func (p *Parser) parseRetStatement() slang.Statement {
	p.pushContext(slang.StatementRETURN.String())

	defer p.context.Pop()

	p.nextToken()

	expr := p.parseExpression(PrecedenceLOWEST)

	return statements.NewReturn(expr)
}

func (p *Parser) parseAssignStatement() slang.Statement {
	ident := expressions.NewIdentifier(p.curToken.Info.Value())

	p.pushContext(slang.StatementASSIGN.String())

	defer p.context.Pop()

	p.nextToken()

	p.nextToken()

	expr := p.parseExpression(PrecedenceLOWEST)

	return statements.NewAssign(ident, expr)
}

func (p *Parser) parseExprStatement() slang.Statement {
	p.pushContext(slang.StatementEXPRESSION.String())

	defer p.context.Pop()

	expr := p.parseExpression(PrecedenceLOWEST)

	return statements.NewExpression(expr)
}
