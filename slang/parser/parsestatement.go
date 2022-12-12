package parser

import (
	"fmt"

	"github.com/jamestunnell/go-synth/slang"
	"github.com/jamestunnell/go-synth/slang/tokens"
)

func (p *Parser) parseStatement() (slang.Statement, error) {
	var s slang.Statement

	var err error

	switch p.curToken.Type() {
	case tokens.TypeIDENT:
		if p.peekToken.Type() == tokens.TypeASSIGN {
			s, err = p.parseAssign()
			if err != nil {
				return nil, fmt.Errorf("failed to parse assign statement: %w", err)
			}
		} else {
			err = NewErrWrongTokenType(tokens.TypeASSIGN, p.peekToken.Type())
		}
	case tokens.TypeRETURN:
		s, err = p.parseReturn()
		if err != nil {
			return nil, fmt.Errorf("failed to parse return statement: %w", err)
		}
	case tokens.TypeIF:
		s, err = p.parseIf()
		if err != nil {
			return nil, fmt.Errorf("failed to parse if statement: %w", err)
		}
	default:
		err = NewErrBadStatementStart(p.curToken)
	}

	return s, err
}
