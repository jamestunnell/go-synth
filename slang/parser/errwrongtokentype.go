package parser

import (
	"fmt"

	"github.com/akrennmair/slice"
	"github.com/jamestunnell/go-synth/slang"
)

type ErrWrongTokenType struct {
	expectedTypes []slang.TokenType
}

func NewErrWrongTokenType(expectedTypes ...slang.TokenType) *ErrWrongTokenType {
	return &ErrWrongTokenType{
		expectedTypes: expectedTypes,
	}
}

func (err *ErrWrongTokenType) Error() string {
	types := slice.Map(err.expectedTypes, func(tokType slang.TokenType) string {
		return tokType.String()
	})

	return fmt.Sprintf("did not find expected token type %s", types)
}
