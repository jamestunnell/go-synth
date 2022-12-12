package parser

import (
	"errors"

	"github.com/jamestunnell/go-synth/slang"
)

type Parser struct {
	Statements []slang.Statement

	lexer slang.Lexer

	curToken  slang.Token
	peekToken slang.Token
}

var (
	errEmptyFuncBody = errors.New("function body is empty")
	errMissingReturn = errors.New("function missing return")
)

func New(l slang.Lexer) *Parser {
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

func (p *Parser) nextTokenSkipLines() {
	p.nextToken()

	for p.curTokenIs(slang.TokenLINE) {
		p.nextToken()
	}
}

func (p *Parser) Run() error {
	statements, err := p.parseStatementsUntil(slang.TokenEOF)
	if err != nil {
		return err
	}

	p.Statements = statements

	return nil
}

func (p *Parser) parseStatementsUntil(stopTokType slang.TokenType) ([]slang.Statement, error) {
	statements := []slang.Statement{}

	for !p.curTokenIs(stopTokType) {
		st, err := p.parseStatement()
		if err != nil {
			return []slang.Statement{}, err
		}

		if st != nil {
			statements = append(statements, st)
		}

		p.nextTokenSkipLines()
	}

	return statements, nil
}

func (p *Parser) curTokenIs(expectedType slang.TokenType) bool {
	return p.curToken.Type() == expectedType
}

func (p *Parser) peekTokenIs(expectedType slang.TokenType) bool {
	return p.peekToken.Type() == expectedType
}

func (p *Parser) curTokenMustBe(expectedType slang.TokenType) error {
	if p.curToken.Type() != expectedType {
		return NewErrWrongTokenType(expectedType, p.curToken.Type())
	}

	return nil
}
