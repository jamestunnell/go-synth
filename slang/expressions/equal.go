package expressions

import "github.com/jamestunnell/go-synth/slang"

type Equal struct {
	*BinaryOperation
}

func NewEqual(left, right slang.Expression) *Equal {
	return &Equal{
		BinaryOperation: NewBinaryOperation(EqualOperator, left, right),
	}
}

func (a *Equal) Equal(other slang.Expression) bool {
	a2, ok := other.(*Equal)
	if !ok {
		return false
	}

	return a.BinaryOperation.Equal(a2.BinaryOperation)
}