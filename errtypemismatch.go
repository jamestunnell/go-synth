package synth

import "fmt"

type ErrTypeMismatch struct {
	Expected, Actual string
}

func NewErrTypeMismatch(expected, actual string) *ErrTypeMismatch {
	return &ErrTypeMismatch{
		Expected: expected,
		Actual:   actual,
	}
}

func (err *ErrTypeMismatch) Error() string {
	const strFmt = "types do not match: expected %s, got %s"

	return fmt.Sprintf(strFmt, err.Expected, err.Actual)
}
