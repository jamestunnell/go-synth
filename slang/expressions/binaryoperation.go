package expressions

import (
	"github.com/rs/zerolog/log"

	"github.com/jamestunnell/go-synth/slang"
)

type BinaryOperator int

type BinaryOperation struct {
	left, right slang.Expression
	operator    BinaryOperator
}

const (
	AddOperator BinaryOperator = iota
	SubtractOperator
	MultiplyOperator
	DivideOperator
	EqualOperator
	NotEqualOperator
	LessOperator
	LessEqualOperator
	GreaterOperator
	GreaterEqualOperator
)

func NewBinaryOperation(op BinaryOperator, left, right slang.Expression) *BinaryOperation {
	return &BinaryOperation{
		operator: op,
		left:     left,
		right:    right,
	}
}

func (bo *BinaryOperation) Equal(other *BinaryOperation) bool {
	if other.operator != bo.operator {
		return false
	}

	return other.left.Equal(bo.left) && other.right.Equal(bo.right)
}

// func (bo *BinaryOperation) String() string {
// 	return fmt.Sprintf("%s %s %s", bo.left, bo.operator, bo.right)
// }

func (operator BinaryOperator) MakeExpression(l, r slang.Expression) slang.Expression {
	var expr slang.Expression
	switch operator {
	case AddOperator:
		expr = NewAdd(l, r)
	case SubtractOperator:
		expr = NewSubtract(l, r)
	case MultiplyOperator:
		expr = NewMultiply(l, r)
	case DivideOperator:
		expr = NewDivide(l, r)
	case EqualOperator:
		expr = NewEqual(l, r)
	case NotEqualOperator:
		expr = NewNotEqual(l, r)
	case LessOperator:
		expr = NewLess(l, r)
	case LessEqualOperator:
		expr = NewLessEqual(l, r)
	case GreaterOperator:
		expr = NewGreater(l, r)
	case GreaterEqualOperator:
		expr = NewGreaterEqual(l, r)
	default:
		log.Fatal().Msgf("unexpected operator %d", operator)
	}

	return expr
}
