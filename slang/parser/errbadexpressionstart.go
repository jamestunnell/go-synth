package parser

import (
	"fmt"

	"github.com/jamestunnell/go-synth/slang"
)

type ErrBadExpressionStart struct {
	StartToken slang.Token
}

func NewErrBadExpressionStart(tok slang.Token) *ErrBadExpressionStart {
	return &ErrBadExpressionStart{
		StartToken: tok,
	}
}

func (err *ErrBadExpressionStart) Error() string {
	const strFmt = "bad expression start: token {type: %s, value: %s}"

	return fmt.Sprintf(strFmt, err.StartToken.Type(), err.StartToken.Value())
}
