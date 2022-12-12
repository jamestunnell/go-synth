package parser

import (
	"github.com/jamestunnell/go-synth/slang"
	"github.com/jamestunnell/go-synth/slang/expressions"
	"github.com/jamestunnell/go-synth/slang/tokens"
)

func AsBinaryOperator(t slang.Token) (expressions.BinaryOperator, bool) {
	switch t.Type() {
	case tokens.TypeSTAR:
		return expressions.MultiplyOperator, true
	case tokens.TypeSLASH:
		return expressions.DivideOperator, true
	case tokens.TypePLUS:
		return expressions.AddOperator, true
	case tokens.TypeMINUS:
		return expressions.SubtractOperator, true
	case tokens.TypeEQUAL:
		return expressions.EqualOperator, true
	case tokens.TypeNOTEQUAL:
		return expressions.NotEqualOperator, true
	case tokens.TypeLESS:
		return expressions.LessOperator, true
	case tokens.TypeLESSEQUAL:
		return expressions.LessEqualOperator, true
	case tokens.TypeGREATER:
		return expressions.GreaterOperator, true
	case tokens.TypeGREATEREQUAL:
		return expressions.GreaterEqualOperator, true
	}

	return 0, false
}
