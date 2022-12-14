package expressions

import "github.com/jamestunnell/go-synth/slang"

type Divide struct {
	*BinaryOperation
}

func NewDivide(left, right slang.Expression) slang.Expression {
	return &Divide{
		BinaryOperation: NewBinaryOperation(DivideOperator, left, right),
	}
}

func (d *Divide) Type() slang.ExprType { return slang.ExprDIVIDE }

func (d *Divide) Equal(other slang.Expression) bool {
	d2, ok := other.(*Divide)
	if !ok {
		return false
	}

	return d.BinaryOperation.Equal(d2.BinaryOperation)
}
