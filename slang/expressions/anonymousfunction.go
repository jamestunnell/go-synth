package expressions

import (
	"github.com/jamestunnell/go-synth/slang"
)

type AnonymousFunction struct {
	ArgNames   []string
	Statements []slang.Statement
}

func NewAnonymousFunction(
	argNames []string, statements []slang.Statement) *AnonymousFunction {
	return &AnonymousFunction{
		ArgNames:   argNames,
		Statements: statements,
	}
}

func (af *AnonymousFunction) Type() slang.ExprType {
	return slang.ExprANONYMOUSFUNC
}

func (af *AnonymousFunction) Equal(other slang.Expression) bool {
	af2, ok := other.(*AnonymousFunction)
	if !ok {
		return false
	}

	if len(af2.Statements) != len(af.Statements) {
		return false
	}

	for i, s := range af.Statements {
		if !s.Equal(af2.Statements[i]) {
			return false
		}
	}

	return true
}
