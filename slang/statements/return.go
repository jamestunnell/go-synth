package statements

import (
	"fmt"

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

func (r *Return) String() string {
	return fmt.Sprintf("return %s", r.value)
}
