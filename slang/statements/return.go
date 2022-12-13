package statements

import (
	"github.com/jamestunnell/go-synth/slang"
)

type Return struct {
	value slang.Expression
}

func NewReturn(value slang.Expression) *Return {
	return &Return{value: value}
}

func (r *Return) Type() slang.StatementType {
	return slang.StatementRETURN
}

func (r *Return) Equal(other slang.Statement) bool {
	r2, ok := other.(*Return)
	if !ok {
		return false
	}

	return r2.value.Equal(r.value)
}
