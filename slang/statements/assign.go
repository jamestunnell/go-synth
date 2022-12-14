package statements

import (
	"github.com/jamestunnell/go-synth/slang"
)

type Assign struct {
	ident slang.Expression
	value slang.Expression
}

func NewAssign(ident slang.Expression, val slang.Expression) slang.Statement {
	return &Assign{
		ident: ident,
		value: val,
	}
}

func (a *Assign) Type() slang.StatementType {
	return slang.StatementASSIGN
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
