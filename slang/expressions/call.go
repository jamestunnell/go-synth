package expressions

import (
	"github.com/jamestunnell/go-synth/slang"
	"golang.org/x/exp/slices"
)

type Call struct {
	Function  slang.Expression // Identifier or FunctionLiteral
	Arguments []slang.Expression
}

func NewCall(fn slang.Expression, args ...slang.Expression) slang.Expression {
	return &Call{
		Function:  fn,
		Arguments: args,
	}
}

func (c *Call) Type() slang.ExprType {
	return slang.ExprCALL
}

func (c *Call) Equal(other slang.Expression) bool {
	c2, ok := other.(*Call)
	if !ok {
		return false
	}

	if !c2.Function.Equal(c.Function) {
		return false
	}

	return slices.EqualFunc(c.Arguments, c2.Arguments, expressionsEqual)
}

func expressionsEqual(a, b slang.Expression) bool {
	return a.Equal(b)
}
