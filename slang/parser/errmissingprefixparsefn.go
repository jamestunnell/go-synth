package parser

import (
	"fmt"

	"github.com/jamestunnell/go-synth/slang"
)

type ErrMissingPrefixParseFn struct {
	Type slang.TokenType
}

func NewErrMissingPrefixParseFn(typ slang.TokenType) *ErrMissingPrefixParseFn {
	return &ErrMissingPrefixParseFn{
		Type: typ,
	}
}

func (err *ErrMissingPrefixParseFn) Error() string {
	const strFmt = "missing prefix parse function for token type %s"

	return fmt.Sprintf(strFmt, err.Type)
}
