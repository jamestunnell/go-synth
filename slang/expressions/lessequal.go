package expressions

import "github.com/jamestunnell/go-synth/slang"

type LessEqual struct {
	*BinaryOperation
}

func NewLessEqual(left, right slang.Expression) slang.Expression {
	return &LessEqual{
		BinaryOperation: NewBinaryOperation(LessEqualOperator, left, right),
	}
}

func (leq *LessEqual) Type() slang.ExprType { return slang.ExprLESSEQUAL }

func (leq *LessEqual) Equal(other slang.Expression) bool {
	leq2, ok := other.(*LessEqual)
	if !ok {
		return false
	}

	return leq.BinaryOperation.Equal(leq2.BinaryOperation)
}
