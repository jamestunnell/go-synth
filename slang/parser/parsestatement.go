package parser

import (
	"errors"

	"github.com/jamestunnell/go-synth/slang"
)

var errBadStatementStart = errors.New("bad statment start")

func (p *Parser) parseStatement() (slang.Statement, *ParseErr) {
	var s slang.Statement

	var pErr *ParseErr

	switch p.curToken.Info.Type() {
	case slang.TokenIDENT:
		s, pErr = p.parseAssign()
	case slang.TokenRETURN:
		s, pErr = p.parseReturn()
	case slang.TokenIF:
		s, pErr = p.parseIf()
	default:
		err := errBadStatementStart

		pErr = NewParseError(err, p.curToken, p.currentContext())
	}

	return s, pErr
}
