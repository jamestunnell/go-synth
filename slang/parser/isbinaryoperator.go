package parser

import (
	"github.com/jamestunnell/go-synth/slang"
	"github.com/jamestunnell/go-synth/slang/expressions"
	"github.com/jamestunnell/go-synth/slang/tokens"
)

func AsBinaryOperator(t slang.Token) (expressions.BinaryOperator, bool) {
	switch t.Type() {
	case tokens.TypeSTAR:
		return expressions.Multiply, true
	case tokens.TypeSLASH:
		return expressions.Divide, true
	case tokens.TypePLUS:
		return expressions.Add, true
	case tokens.TypeMINUS:
		return expressions.Subtract, true
	}

	return 0, false
}
