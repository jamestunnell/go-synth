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

func (r *Assign) Type() string {
	return TypeASSIGN
}
