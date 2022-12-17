package expressions

import (
	"github.com/jamestunnell/go-synth/slang"
)

type Float struct {
	Value float64
}

func NewFloat(val float64) *Float {
	return &Float{Value: val}
}

func (f *Float) Type() slang.ExprType { return slang.ExprFLOAT }

func (f *Float) Equal(other slang.Expression) bool {
	f2, ok := other.(*Float)
	if !ok {
		return false
	}

	return f2.Value == f.Value
}
