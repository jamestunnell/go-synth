package statements

import (
	"github.com/jamestunnell/go-synth/slang"
)

type Expression struct {
	value slang.Expression
}

func NewExpression(value slang.Expression) *Expression {
	return &Expression{value: value}
}

func (e *Expression) Type() slang.StatementType {
	return slang.StatementEXPRESSION
}

func (e *Expression) Equal(other slang.Statement) bool {
	e2, ok := other.(*Expression)
	if !ok {
		return false
	}

	return e2.value.Equal(e.value)
}
