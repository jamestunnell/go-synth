package expressions

import "github.com/jamestunnell/go-synth/slang"

type GreaterEqual struct {
	*BinaryOperation
}

func NewGreaterEqual(left, right slang.Expression) slang.Expression {
	return &GreaterEqual{
		BinaryOperation: NewBinaryOperation(GreaterEqualOperator, left, right),
	}
}

func (geq *GreaterEqual) Type() slang.ExprType { return slang.ExprGREATEREQUAL }

func (geq *GreaterEqual) Equal(other slang.Expression) bool {
	geq2, ok := other.(*GreaterEqual)
	if !ok {
		return false
	}

	return geq.BinaryOperation.Equal(geq2.BinaryOperation)
}
