package statements

import (
	"fmt"

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

func (a *Assign) String() string {
	return fmt.Sprintf("%s = %s", a.ident.Name, a.value.String())
}
