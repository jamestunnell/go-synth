package parser

import (
	"fmt"

	"github.com/jamestunnell/go-synth/slang"
)

type ErrWrongTokenType struct {
	Expected, Actual slang.TokenType
}

func NewErrWrongTokenType(expected, actual slang.TokenType) *ErrWrongTokenType {
	return &ErrWrongTokenType{
		Expected: expected,
		Actual:   actual,
	}
}

func (err *ErrWrongTokenType) Error() string {
	return fmt.Sprintf("wrong token type: expected %s, got %s", err.Expected, err.Actual)
}
