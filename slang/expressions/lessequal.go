package expressions

import "github.com/jamestunnell/go-synth/slang"

type LessEqual struct {
	*BinaryOperation
}

func NewLessEqual(left, right slang.Expression) *LessEqual {
	return &LessEqual{
		BinaryOperation: NewBinaryOperation(LessEqualOperator, left, right),
	}
}

func (a *LessEqual) Equal(other slang.Expression) bool {
	a2, ok := other.(*LessEqual)
	if !ok {
		return false
	}

	return a.BinaryOperation.Equal(a2.BinaryOperation)
}
