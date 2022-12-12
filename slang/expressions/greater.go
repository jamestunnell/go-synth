package expressions

import "github.com/jamestunnell/go-synth/slang"

type Greater struct {
	*BinaryOperation
}

func NewGreater(left, right slang.Expression) *Greater {
	return &Greater{
		BinaryOperation: NewBinaryOperation(GreaterOperator, left, right),
	}
}

func (a *Greater) Equal(other slang.Expression) bool {
	a2, ok := other.(*Greater)
	if !ok {
		return false
	}

	return a.BinaryOperation.Equal(a2.BinaryOperation)
}
