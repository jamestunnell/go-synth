package parser

import (
	"fmt"
)

type ErrNotImplemented struct {
	Type string
}

func NewErrNotImplemented(typ string) *ErrNotImplemented {
	return &ErrNotImplemented{
		Type: typ,
	}
}

func (err *ErrNotImplemented) Error() string {
	const strFmt = "support for %s is not implemented"

	return fmt.Sprintf(strFmt, err.Type)
}
