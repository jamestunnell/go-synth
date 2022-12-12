package parser

import (
	"errors"

	"github.com/jamestunnell/go-synth/slang"
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

func (p *Parser) nextTokenSkipLines() {
	p.nextToken()

	for p.curToken.Type() == tokens.TypeLINE {
		p.nextToken()
	}
}

func (p *Parser) ParseProgram() (*slang.Program, error) {
	program := slang.NewProgram()

	statements, err := p.parseStatementsUntil(tokens.TypeEOF)
	if err != nil {
		return nil, err
	}

	program.Statements = statements

	return program, nil
}

func (p *Parser) parseStatementsUntil(stopTokType string) ([]slang.Statement, error) {
	statements := []slang.Statement{}

	for p.curToken.Type() != stopTokType {
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

func (p *Parser) curTokenIs(expectedType string) bool {
	return p.curToken.Type() == expectedType
}

func (p *Parser) peekTokenIs(expectedType string) bool {
	return p.peekToken.Type() == expectedType
}

func (p *Parser) curTokenMustBe(expectedType string) error {
	if p.curToken.Type() != expectedType {
		return NewErrWrongTokenType(expectedType, p.curToken.Type())
	}

	return nil
}
