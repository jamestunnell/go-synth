package synth

import "fmt"

type ErrUnsupportedType struct {
	Val any
}

func NewErrUnsupportedType(val any) *ErrUnsupportedType {
	return &ErrUnsupportedType{Val: val}
}

func (err *ErrUnsupportedType) Error() string {
	return fmt.Sprintf("value %v type is not unsupported", err.Val)
}
