package parser

import "fmt"

type ErrWrongTokenType struct {
	Expected, Actual string
}

func NewErrWrongTokenType(expected, actual string) *ErrWrongTokenType {
	return &ErrWrongTokenType{
		Expected: expected,
		Actual:   actual,
	}
}

func (err *ErrWrongTokenType) Error() string {
	return fmt.Sprintf("wrong token type: expected %s, got %s", err.Expected, err.Actual)
}
