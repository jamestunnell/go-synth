package parser

import "github.com/jamestunnell/go-synth/slang"

func (p *Parser) parseGroupedExpr() (slang.Expression, *ParseErr) {
	err := NewErrNotImplemented("grouped expression")

	return nil, NewParseError(err, p.curToken, p.currentContext())
}
