package expressions

import "github.com/jamestunnell/go-synth/slang"

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
