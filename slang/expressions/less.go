package expressions

import "github.com/jamestunnell/go-synth/slang"

type Less struct {
	*BinaryOperation
}

func NewLess(left, right slang.Expression) *Less {
	return &Less{
		BinaryOperation: NewBinaryOperation(LessOperator, left, right),
	}
}

func (l *Less) Type() slang.ExprType { return slang.ExprLESS }

func (l *Less) Equal(other slang.Expression) bool {
	l2, ok := other.(*Less)
	if !ok {
		return false
	}

	return l.BinaryOperation.Equal(l2.BinaryOperation)
}
