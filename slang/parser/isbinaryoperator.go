package parser

import (
	"github.com/jamestunnell/go-synth/slang"
	"github.com/jamestunnell/go-synth/slang/expressions"
)

func AsBinaryOperator(tt slang.TokenType) (expressions.BinaryOperator, Precedence, bool) {
	switch tt {
	case slang.TokenSTAR:
		return expressions.MultiplyOperator, PrecedencePRODUCT, true
	case slang.TokenSLASH:
		return expressions.DivideOperator, PrecedencePRODUCT, true
	case slang.TokenPLUS:
		return expressions.AddOperator, PrecedenceSUM, true
	case slang.TokenMINUS:
		return expressions.SubtractOperator, PrecedenceSUM, true
	case slang.TokenEQUAL:
		return expressions.EqualOperator, PrecedenceEQUALS, true
	case slang.TokenNOTEQUAL:
		return expressions.NotEqualOperator, PrecedenceEQUALS, true
	case slang.TokenLESS:
		return expressions.LessOperator, PrecedenceLESSGREATER, true
	case slang.TokenLESSEQUAL:
		return expressions.LessEqualOperator, PrecedenceLESSGREATER, true
	case slang.TokenGREATER:
		return expressions.GreaterOperator, PrecedenceLESSGREATER, true
	case slang.TokenGREATEREQUAL:
		return expressions.GreaterEqualOperator, PrecedenceLESSGREATER, true
	}

	return 0, PrecedenceLOWEST, false
}
