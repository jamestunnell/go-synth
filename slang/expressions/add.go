package expressions

import "github.com/jamestunnell/go-synth/slang"

type Add struct {
	*BinaryOperation
}

func NewAdd(left, right slang.Expression) *Add {
	return &Add{
		BinaryOperation: NewBinaryOperation(AddOperator, left, right),
	}
}

func (a *Add) Type() slang.ExprType { return slang.ExprADD }

func (a *Add) Equal(other slang.Expression) bool {
	a2, ok := other.(*Add)
	if !ok {
		return false
	}

	return a.BinaryOperation.Equal(a2.BinaryOperation)
}
