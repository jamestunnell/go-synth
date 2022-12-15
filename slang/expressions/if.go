package expressions

import (
	"github.com/jamestunnell/go-synth/slang"
	"github.com/jamestunnell/go-synth/slang/statements"
)

type If struct {
	Condition   slang.Expression
	Consequence *statements.Block
}

func NewIf(cond slang.Expression, conseq *statements.Block) *If {
	return &If{
		Condition:   cond,
		Consequence: conseq,
	}
}

func (i *If) Type() slang.ExprType { return slang.ExprIF }

func (i *If) Equal(other slang.Expression) bool {
	i2, ok := other.(*If)
	if !ok {
		return false
	}

	return i2.Condition.Equal(i.Condition) && i2.Consequence.Equal(i.Consequence)
}
