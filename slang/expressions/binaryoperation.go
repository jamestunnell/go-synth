package expressions

import (
	"fmt"

	"github.com/jamestunnell/go-synth/slang"
)

type BinaryOperator int

type BinaryOperation struct {
	left, right slang.Expression
	operator    BinaryOperator
}

const (
	Add BinaryOperator = iota
	Subtract
	Multiply
	Divide
)

func NewBinaryOperation(op BinaryOperator, left, right slang.Expression) *BinaryOperation {
	return &BinaryOperation{
		operator: op,
		left:     left,
		right:    right,
	}
}

func (bo *BinaryOperation) String() string {
	return fmt.Sprintf("%s %s %s", bo.left, bo.operator, bo.right)
}

func (bo BinaryOperator) String() string {
	var str string

	switch bo {
	case Add:
		str = "+"
	case Subtract:
		str = "-"
	case Multiply:
		str = "*"
	case Divide:
		str = "/"
	}

	return str
}
