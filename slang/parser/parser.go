package parser

import (
	"errors"

	"github.com/jamestunnell/go-synth/slang"
	"github.com/jamestunnell/go-synth/util/stack"
)

type prefixParseFn func() slang.Expression
type infixParseFn func(slang.Expression) slang.Expression

type Parser struct {
	Statements []slang.Statement
	Errors     []*ParseErr

	lexer   slang.Lexer
	context *stack.Stack[*ParseContext]

	curToken  *slang.Token
	peekToken *slang.Token

	prefixParseFns map[slang.TokenType]prefixParseFn
	infixParseFns  map[slang.TokenType]infixParseFn
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
		Statements:     []slang.Statement{},
		Errors:         []*ParseErr{},
		context:        stack.New[*ParseContext](),
		lexer:          l,
		prefixParseFns: map[slang.TokenType]prefixParseFn{},
		infixParseFns:  map[slang.TokenType]infixParseFn{},
	}

	p.registerPrefix(slang.TokenIDENT, p.parseIdentifier)
	p.registerPrefix(slang.TokenINT, p.parseInteger)
	p.registerPrefix(slang.TokenFLOAT, p.parseFloat)
	p.registerPrefix(slang.TokenTRUE, p.parseTrue)
	p.registerPrefix(slang.TokenFALSE, p.parseFalse)
	p.registerPrefix(slang.TokenMINUS, p.parseNegative)
	p.registerPrefix(slang.TokenBANG, p.parseNot)
	p.registerPrefix(slang.TokenLPAREN, p.parseGroupedExpression)
	p.registerPrefix(slang.TokenIF, p.parseIfExpression)
	p.registerPrefix(slang.TokenFUNC, p.parseFuncLiteral)

	p.registerInfix(slang.TokenPLUS, p.parseAdd)
	p.registerInfix(slang.TokenMINUS, p.parseSubtract)
	p.registerInfix(slang.TokenSTAR, p.parseMultiply)
	p.registerInfix(slang.TokenSLASH, p.parseDivide)
	p.registerInfix(slang.TokenEQUAL, p.parseEqual)
	p.registerInfix(slang.TokenNOTEQUAL, p.parseNotEqual)
	p.registerInfix(slang.TokenLESS, p.parseLess)
	p.registerInfix(slang.TokenGREATER, p.parseGreater)

	// Read two tokens, so curToken and peekToken are both set
	p.nextToken()

	p.nextToken()

	return p
}
func (p *Parser) registerPrefix(tokenType slang.TokenType, fn prefixParseFn) {
	p.prefixParseFns[tokenType] = fn
}

func (p *Parser) registerInfix(tokenType slang.TokenType, fn infixParseFn) {
	p.infixParseFns[tokenType] = fn
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
		if st := p.parseStatement(); st != nil {
			statements = append(statements, st)
		}

		p.nextToken()
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

// func (p *Parser) curTokenMustBe(expectedType slang.TokenType) bool {
// 	if p.curToken.Info.Type() == expectedType {
// 		return true
// 	}

// 	err := NewErrWrongTokenType(expectedType)
// 	pErr := NewParseError(err, p.curToken, p.currentContext())

// 	p.Errors = append(p.Errors, pErr)

// 	return false
// }

func (p *Parser) expectPeek(expectedType slang.TokenType) bool {
	if p.peekToken.Info.Type() != expectedType {
		p.peekError(expectedType)

		return false
	}

	p.nextToken()

	return true
}

func (p *Parser) peekError(expectedType slang.TokenType) {
	err := NewErrWrongTokenType(expectedType)
	pErr := NewParseError(err, p.peekToken, p.currentContext())

	p.Errors = append(p.Errors, pErr)
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

func (p *Parser) peekPrecedence() Precedence {
	if p, ok := precedences[p.peekToken.Info.Type()]; ok {
		return p
	}

	return PrecedenceLOWEST
}

func (p *Parser) curPrecedence() Precedence {
	if p, ok := precedences[p.curToken.Info.Type()]; ok {
		return p
	}

	return PrecedenceLOWEST
}
