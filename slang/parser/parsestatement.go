package parser

import (
	"fmt"

	"github.com/jamestunnell/go-synth/slang"
)

func (p *Parser) parseStatement() (slang.Statement, error) {
	var s slang.Statement

	var err error

	switch p.curToken.Type() {
	case slang.TokenIDENT:
		if p.peekTokenIs(slang.TokenASSIGN) {
			s, err = p.parseAssign()
			if err != nil {
				return nil, fmt.Errorf("failed to parse assign statement: %w", err)
			}
		} else {
			err = NewErrWrongTokenType(slang.TokenASSIGN, p.peekToken.Type())
		}
	case slang.TokenRETURN:
		s, err = p.parseReturn()
		if err != nil {
			return nil, fmt.Errorf("failed to parse return statement: %w", err)
		}
	case slang.TokenIF:
		s, err = p.parseIf()
		if err != nil {
			return nil, fmt.Errorf("failed to parse if statement: %w", err)
		}
	default:
		err = NewErrBadStatementStart(p.curToken)
	}

	return s, err
}
