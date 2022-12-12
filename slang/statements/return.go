package statements

import (
	"github.com/jamestunnell/go-synth/slang"
)

type Return struct {
	value slang.Expression
}

const TypeRETURN = "RETURN"

func NewReturn(value slang.Expression) *Return {
	return &Return{value: value}
}

func (r *Return) Type() string {
	return TypeRETURN
}

func (r *Return) Equal(other slang.Statement) bool {
	r2, ok := other.(*Return)
	if !ok {
		return false
	}

	return r2.value.Equal(r.value)
}
