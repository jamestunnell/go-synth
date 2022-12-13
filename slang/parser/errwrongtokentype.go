package parser

import (
	"fmt"

	"github.com/jamestunnell/go-synth/slang"
)

type ErrWrongTokenType struct {
	expectedType slang.TokenType
}

func NewErrWrongTokenType(expectedType slang.TokenType) *ErrWrongTokenType {
	return &ErrWrongTokenType{
		expectedType: expectedType,
	}
}

func (err *ErrWrongTokenType) Error() string {
	return fmt.Sprintf("did not find expected token type %s", err.expectedType)
}
