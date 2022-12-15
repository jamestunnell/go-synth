package expressions

import (
	"github.com/jamestunnell/go-synth/slang"
	"github.com/jamestunnell/go-synth/slang/statements"
)

type IfElse struct {
	Condition                slang.Expression
	Consequence, Alternative *statements.Block
}

func NewIfElse(cond slang.Expression, conseq, altern *statements.Block) *IfElse {
	return &IfElse{
		Condition:   cond,
		Consequence: conseq,
		Alternative: altern,
	}
}

func (i *IfElse) Type() slang.ExprType { return slang.ExprIF }

func (i *IfElse) Equal(other slang.Expression) bool {
	i2, ok := other.(*IfElse)
	if !ok {
		return false
	}

	return i2.Condition.Equal(i.Condition) &&
		i2.Consequence.Equal(i.Consequence) &&
		i2.Alternative.Equal(i.Alternative)
}
