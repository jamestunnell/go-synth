package expressions

import "github.com/jamestunnell/go-synth/slang"

type Subtract struct {
	*BinaryOperation
}

func NewSubtract(left, right slang.Expression) *Subtract {
	return &Subtract{
		BinaryOperation: NewBinaryOperation(SubtractOperator, left, right),
	}
}

func (a *Subtract) Equal(other slang.Expression) bool {
	a2, ok := other.(*Subtract)
	if !ok {
		return false
	}

	return a.BinaryOperation.Equal(a2.BinaryOperation)
}
