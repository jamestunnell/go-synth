package parser

import "github.com/jamestunnell/go-synth/slang"

func (p *Parser) parseGroupedExpr() (slang.Expression, *ParseErr) {
	err := NewErrNotImplemented("grouped expression")
	pErr := NewParseError(err, p.curToken, p.currentContext())

	p.skipToNextLineOrEOF()

	return nil, pErr
}
