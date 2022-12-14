package expressions

import "github.com/jamestunnell/go-synth/slang"

type Multiply struct {
	*BinaryOperation
}

func NewMultiply(left, right slang.Expression) slang.Expression {
	return &Multiply{
		BinaryOperation: NewBinaryOperation(MultiplyOperator, left, right),
	}
}

func (m *Multiply) Type() slang.ExprType { return slang.ExprMULTIPLY }

func (m *Multiply) Equal(other slang.Expression) bool {
	m2, ok := other.(*Multiply)
	if !ok {
		return false
	}

	return m.BinaryOperation.Equal(m2.BinaryOperation)
}
