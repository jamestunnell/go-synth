package expressions

import (
	"fmt"
	"strings"

	"github.com/akrennmair/slice"
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

func (af *AnonymousFunction) String() string {
	argsStr := strings.Join(af.ArgNames, ", ")

	statementToStr := func(s slang.Statement) string {
		return s.String()
	}
	statementsStr := slice.Map(af.Statements, statementToStr)

	return fmt.Sprintf("func(%s) {\n%s\n}", argsStr, statementsStr)
}
