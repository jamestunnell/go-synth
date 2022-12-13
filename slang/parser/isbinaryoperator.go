package parser

import (
	"github.com/jamestunnell/go-synth/slang"
	"github.com/jamestunnell/go-synth/slang/expressions"
)

func AsBinaryOperator(tt slang.TokenType) (expressions.BinaryOperator, bool) {
	switch tt {
	case slang.TokenSTAR:
		return expressions.MultiplyOperator, true
	case slang.TokenSLASH:
		return expressions.DivideOperator, true
	case slang.TokenPLUS:
		return expressions.AddOperator, true
	case slang.TokenMINUS:
		return expressions.SubtractOperator, true
	case slang.TokenEQUAL:
		return expressions.EqualOperator, true
	case slang.TokenNOTEQUAL:
		return expressions.NotEqualOperator, true
	case slang.TokenLESS:
		return expressions.LessOperator, true
	case slang.TokenLESSEQUAL:
		return expressions.LessEqualOperator, true
	case slang.TokenGREATER:
		return expressions.GreaterOperator, true
	case slang.TokenGREATEREQUAL:
		return expressions.GreaterEqualOperator, true
	}

	return 0, false
}
