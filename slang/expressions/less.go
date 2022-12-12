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

func (a *Less) Equal(other slang.Expression) bool {
	a2, ok := other.(*Less)
	if !ok {
		return false
	}

	return a.BinaryOperation.Equal(a2.BinaryOperation)
}
