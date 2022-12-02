package parser

import (
	"fmt"

	"github.com/jamestunnell/go-synth/slang"
)

type ErrBadStatementStart struct {
	StartToken slang.Token
}

func NewErrBadStatementStart(tok slang.Token) *ErrBadStatementStart {
	return &ErrBadStatementStart{
		StartToken: tok,
	}
}

func (err *ErrBadStatementStart) Error() string {
	const strFmt = "bad statement start: token {type: %s, value: %s}"

	return fmt.Sprintf(strFmt, err.StartToken.Type(), err.StartToken.Value())
}
