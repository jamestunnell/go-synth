package expressions

import "github.com/jamestunnell/go-synth/slang"

type GreaterEqual struct {
	*BinaryOperation
}

func NewGreaterEqual(left, right slang.Expression) *GreaterEqual {
	return &GreaterEqual{
		BinaryOperation: NewBinaryOperation(GreaterEqualOperator, left, right),
	}
}

func (a *GreaterEqual) Equal(other slang.Expression) bool {
	a2, ok := other.(*GreaterEqual)
	if !ok {
		return false
	}

	return a.BinaryOperation.Equal(a2.BinaryOperation)
}
