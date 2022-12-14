package expressions

import "github.com/jamestunnell/go-synth/slang"

type Bool struct {
	Value bool
}

func NewBool(val bool) *Bool {
	return &Bool{Value: val}
}

func (b *Bool) Type() slang.ExprType { return slang.ExprBOOL }

func (b *Bool) Equal(other slang.Expression) bool {
	b2, ok := other.(*Bool)
	if !ok {
		return false
	}

	return b2.Value == b.Value
}
