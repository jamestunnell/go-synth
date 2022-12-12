package expressions

import "github.com/jamestunnell/go-synth/slang"

type Multiply struct {
	*BinaryOperation
}

func NewMultiply(left, right slang.Expression) *Multiply {
	return &Multiply{
		BinaryOperation: NewBinaryOperation(MultiplyOperator, left, right),
	}
}

func (a *Multiply) Equal(other slang.Expression) bool {
	a2, ok := other.(*Multiply)
	if !ok {
		return false
	}

	return a.BinaryOperation.Equal(a2.BinaryOperation)
}
