package network

import (
	"fmt"
	"strings"
)

type ErrUntargetedInputs struct {
	inputNames []string
}

func NewErrUntargetedInputs(inputNames []string) *ErrUntargetedInputs {
	return &ErrUntargetedInputs{
		inputNames: inputNames,
	}
}

func (err *ErrUntargetedInputs) Error() string {
	const strFmt = "untargeted input(s): %s"

	namesStr := strings.Join(err.inputNames, ", ")

	return fmt.Sprintf(strFmt, namesStr)
}
