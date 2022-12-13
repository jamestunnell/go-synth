package statements

import (
	"github.com/jamestunnell/go-synth/slang"
)

type If struct {
	Condition slang.Expression
	Body      []slang.Statement
}

func NewIf(condition slang.Expression, body ...slang.Statement) slang.Statement {
	return &If{
		Condition: condition,
		Body:      body,
	}
}

func (i *If) Type() slang.StatementType {
	return slang.StatementIF
}

func (i *If) Equal(other slang.Statement) bool {
	i2, ok := other.(*If)
	if !ok {
		return false
	}

	if len(i2.Body) != len(i.Body) {
		return false
	}

	for i, s := range i.Body {
		if !s.Equal(i2.Body[i]) {
			return false
		}
	}

	return true
}
