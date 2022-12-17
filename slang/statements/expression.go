package statements

import (
	"github.com/jamestunnell/go-synth/slang"
)

type Expression struct {
	Value slang.Expression
}

func NewExpression(val slang.Expression) *Expression {
	return &Expression{Value: val}
}

func (e *Expression) Type() slang.StatementType {
	return slang.StatementEXPRESSION
}

func (e *Expression) Equal(other slang.Statement) bool {
	e2, ok := other.(*Expression)
	if !ok {
		return false
	}

	return e2.Value.Equal(e.Value)
}
