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

func (g *Greater) Type() slang.ExprType { return slang.ExprGREATER }

func (g *Greater) Equal(other slang.Expression) bool {
	g2, ok := other.(*Greater)
	if !ok {
		return false
	}

	return g.BinaryOperation.Equal(g2.BinaryOperation)
}
