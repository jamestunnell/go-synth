package expressions

import "github.com/jamestunnell/go-synth/slang"

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
