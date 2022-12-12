package parser

import (
	"github.com/jamestunnell/go-synth/slang/expressions"
	"github.com/jamestunnell/go-synth/slang/tokens"
)

func (p *Parser) parseIdentifier() (*expressions.Identifier, error) {
	if p.curToken.Type() != tokens.TypeIDENT {
		return nil, NewErrWrongTokenType(tokens.TypeIDENT, p.curToken.Type())
	}

	return expressions.NewIdentifier(p.curToken.Value()), nil
}
