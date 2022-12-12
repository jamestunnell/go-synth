package statements

import (
	"github.com/jamestunnell/go-synth/slang"
	"github.com/jamestunnell/go-synth/slang/expressions"
)

type Assign struct {
	ident *expressions.Identifier
	value slang.Expression
}

const TypeASSIGN = "ASSIGN"

func NewAssign(ident *expressions.Identifier, val slang.Expression) slang.Statement {
	return &Assign{
		ident: ident,
		value: val,
	}
}

func (a *Assign) Type() string {
	return TypeASSIGN
}

func (a *Assign) Equal(other slang.Statement) bool {
	a2, ok := other.(*Assign)
	if !ok {
		return false
	}

	if !a2.ident.Equal(a.ident) {
		return false
	}

	return a2.value.Equal(a.value)
}
