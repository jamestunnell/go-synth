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

		if p.curToken.Type() == tokens.TypeLINE {
			p.nextToken()
		}
	}

	return program, nil
}
