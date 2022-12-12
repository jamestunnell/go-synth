package expressions

import "github.com/jamestunnell/go-synth/slang"

type Divide struct {
	*BinaryOperation
}

func NewDivide(left, right slang.Expression) *Divide {
	return &Divide{
		BinaryOperation: NewBinaryOperation(DivideOperator, left, right),
	}
}

func (a *Divide) Equal(other slang.Expression) bool {
	a2, ok := other.(*Divide)
	if !ok {
		return false
	}

	return a.BinaryOperation.Equal(a2.BinaryOperation)
}
