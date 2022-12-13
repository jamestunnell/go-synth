package parser

import (
	"errors"

	"github.com/jamestunnell/go-synth/slang"
	"github.com/jamestunnell/go-synth/util/stack"
)

type Parser struct {
	Statements []slang.Statement
	Errors     []*ParseErr

	lexer   slang.Lexer
	context *stack.Stack[*ParseContext]

	curToken  *slang.Token
	peekToken *slang.Token
}

type ParseResults struct {
	Statements []slang.Statement
	Errors     []*ParseErr
}

var (
	errEmptyFuncBody = errors.New("function body is empty")
	errMissingReturn = errors.New("function missing return")
)

func New(l slang.Lexer) *Parser {
	p := &Parser{
		Statements: []slang.Statement{},
		Errors:     []*ParseErr{},
		context:    stack.New[*ParseContext](),
		lexer:      l,
	}

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

func (p *Parser) skipToNextLineOrEOF() {
	for !p.curTokenIs(slang.TokenEOF) && !p.curTokenIs(slang.TokenLINE) {
		p.nextToken()
	}

	if p.curTokenIs(slang.TokenLINE) {
		p.nextToken()
	}
}

func (p *Parser) Run() *ParseResults {
	p.pushContext("global")

	defer p.context.Pop()

	statments := p.parseStatementsUntil(slang.TokenEOF)

	p.Statements = append(p.Statements, statments...)

	return &ParseResults{
		Statements: p.Statements,
		Errors:     p.Errors,
	}
}

func (p *Parser) parseStatementsUntil(
	stopTokType slang.TokenType) []slang.Statement {
	statements := []slang.Statement{}

	for !p.curTokenIs(slang.TokenEOF) && !p.curTokenIs(stopTokType) {
		st, err := p.parseStatement()
		if err != nil {
			p.Errors = append(p.Errors, err)
		}

		if st != nil {
			statements = append(statements, st)
		}

		if err != nil {
			p.skipToNextLineOrEOF()

			continue
		}

		p.nextTokenSkipLines()
	}

	// did we stop because of EOF or the expected stop token?
	if !p.curTokenIs(stopTokType) {
		err := NewParseError(
			NewErrWrongTokenType(stopTokType), p.curToken, p.currentContext())

		p.Errors = append(p.Errors, err)
	}

	return statements
}

func (p *Parser) curTokenIs(expectedType slang.TokenType) bool {
	return p.curToken.Info.Type() == expectedType
}

func (p *Parser) peekTokenIs(expectedType slang.TokenType) bool {
	return p.peekToken.Info.Type() == expectedType
}

func (p *Parser) curTokenMustBe(expectedType slang.TokenType) *ParseErr {
	if p.curToken.Info.Type() != expectedType {
		err := NewErrWrongTokenType(expectedType)

		return NewParseError(err, p.curToken, p.currentContext())
	}

	return nil
}

func (p *Parser) peekTokenMustBe(expectedType slang.TokenType) *ParseErr {
	if p.peekToken.Info.Type() != expectedType {
		err := NewErrWrongTokenType(expectedType)

		return NewParseError(err, p.peekToken, p.currentContext())
	}

	return nil
}

func (p *Parser) pushContext(descr string) {
	c := &ParseContext{
		Start:       p.curToken,
		Description: descr,
	}

	p.context.Push(c)
}

func (p *Parser) currentContext() *ParseContext {
	return p.context.Top()
}

func (p *Parser) NewParseErr(err error) *ParseErr {
	return NewParseError(err, p.curToken, p.currentContext())
}
