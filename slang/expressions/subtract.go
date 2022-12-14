package expressions

import "github.com/jamestunnell/go-synth/slang"

type Subtract struct {
	*BinaryOperation
}

func NewSubtract(left, right slang.Expression) slang.Expression {
	return &Subtract{
		BinaryOperation: NewBinaryOperation(SubtractOperator, left, right),
	}
}

func (a *Subtract) Type() slang.ExprType { return slang.ExprSUBTRACT }

func (a *Subtract) Equal(other slang.Expression) bool {
	a2, ok := other.(*Subtract)
	if !ok {
		return false
	}

	return a.BinaryOperation.Equal(a2.BinaryOperation)
}
