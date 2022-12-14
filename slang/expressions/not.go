package expressions

import "github.com/jamestunnell/go-synth/slang"

type Not struct {
	Value slang.Expression
}

func NewNot(val slang.Expression) *Not {
	return &Not{Value: val}
}

func (b *Not) Type() slang.ExprType { return slang.ExprNOT }

func (b *Not) Equal(other slang.Expression) bool {
	b2, ok := other.(*Not)
	if !ok {
		return false
	}

	return b2.Value.Equal(b.Value)
}
