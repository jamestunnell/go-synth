package expressions

import "github.com/jamestunnell/go-synth/slang"

type NotEqual struct {
	*BinaryOperation
}

func NewNotEqual(left, right slang.Expression) *NotEqual {
	return &NotEqual{
		BinaryOperation: NewBinaryOperation(NotEqualOperator, left, right),
	}
}

func (a *NotEqual) Equal(other slang.Expression) bool {
	a2, ok := other.(*NotEqual)
	if !ok {
		return false
	}

	return a.BinaryOperation.Equal(a2.BinaryOperation)
}
