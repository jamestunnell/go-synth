package interpreter

import (
	"github.com/jamestunnell/go-synth/slang"
	"github.com/jamestunnell/go-synth/slang/expressions"
	"github.com/jamestunnell/go-synth/slang/objects"
	"github.com/jamestunnell/go-synth/slang/statements"
	"github.com/rs/zerolog/log"
)

var (
	NULL  = &objects.Null{}
	TRUE  = objects.NewBool(true)
	FALSE = objects.NewBool(false)
)

func EvalStatement(s slang.Statement) slang.Object {
	switch ss := s.(type) {
	case *statements.Expression:
		return EvalExpression(ss.Value)
	}

	return NULL
}

func EvalExpression(e slang.Expression) slang.Object {
	switch ee := e.(type) {
	case *expressions.Integer:
		return objects.NewInteger(ee.Value)
	case *expressions.Float:
		return objects.NewFloat(ee.Value)
	case *expressions.Bool:
		if ee.Value {
			return TRUE
		}

		return FALSE
	case *expressions.Not:
		return evalNot(ee)
	case *expressions.Negative:
		return evalNeg(ee)
	case *expressions.Add:
		return evalBinOp(ee.BinaryOperation)
	case *expressions.Subtract:
		return evalBinOp(ee.BinaryOperation)
	case *expressions.Multiply:
		return evalBinOp(ee.BinaryOperation)
	case *expressions.Divide:
		return evalBinOp(ee.BinaryOperation)
	case *expressions.Equal:
		return evalBinOp(ee.BinaryOperation)
	case *expressions.GreaterEqual:
		return evalBinOp(ee.BinaryOperation)
	case *expressions.LessEqual:
		return evalBinOp(ee.BinaryOperation)
	case *expressions.Less:
		return evalBinOp(ee.BinaryOperation)
	case *expressions.Greater:
		return evalBinOp(ee.BinaryOperation)
	case *expressions.NotEqual:
		return evalBinOp(ee.BinaryOperation)
	}

	return NULL
}

func evalNot(e *expressions.Not) slang.Object {
	subject := EvalExpression(e.Value)
	switch subject {
	case TRUE:
		return FALSE
	case NULL, FALSE:
		return TRUE
	}

	return FALSE
}

func evalNeg(e *expressions.Negative) slang.Object {
	subject := EvalExpression(e.Value)

	switch obj := subject.(type) {
	case *objects.Integer:
		return objects.NewInteger(-obj.Value)
	case *objects.Float:
		return objects.NewFloat(-obj.Value)
	}

	return NULL
}

func evalBinOp(e *expressions.BinaryOperation) slang.Object {
	l := EvalExpression(e.Left)
	r := EvalExpression(e.Right)

	if l.Type() == slang.ObjectINTEGER {
		a := l.(*objects.Integer).Value

		switch r.Type() {
		case slang.ObjectINTEGER:
			b := r.(*objects.Integer).Value

			return makeBinOpResult(a, b, e.Operator, objects.NewInteger)
		case slang.ObjectFLOAT:
			b := r.(*objects.Float).Value

			return makeBinOpResult(float64(a), b, e.Operator, objects.NewFloat)
		}
	}

	if l.Type() == slang.ObjectFLOAT {
		a := l.(*objects.Float).Value

		switch r.Type() {
		case slang.ObjectFLOAT:
			b := r.(*objects.Float).Value

			return makeBinOpResult(a, b, e.Operator, objects.NewFloat)
		case slang.ObjectINTEGER:
			b := r.(*objects.Integer).Value

			return makeBinOpResult(a, float64(b), e.Operator, objects.NewFloat)
		}
	}

	if l.Type() == slang.ObjectINTEGER && r.Type() == slang.ObjectINTEGER {
		a := l.(*objects.Integer).Value
		b := r.(*objects.Integer).Value

		return makeBinOpResult(a, b, e.Operator, objects.NewInteger)
	}

	return NULL
}

func makeBinOpResult[T int64 | float64](
	a, b T,
	operator expressions.BinaryOperator,
	fn func(T) slang.Object) slang.Object {
	var obj slang.Object

	switch operator {
	case expressions.AddOperator:
		obj = fn(a + b)
	case expressions.SubtractOperator:
		obj = fn(a - b)
	case expressions.MultiplyOperator:
		obj = fn(a * b)
	case expressions.DivideOperator:
		obj = fn(a / b)
	case expressions.EqualOperator:
		obj = objects.NewBool(a == b)
	case expressions.NotEqualOperator:
		obj = objects.NewBool(a != b)
	case expressions.LessOperator:
		obj = objects.NewBool(a < b)
	case expressions.LessEqualOperator:
		obj = objects.NewBool(a <= b)
	case expressions.GreaterOperator:
		obj = objects.NewBool(a > b)
	case expressions.GreaterEqualOperator:
		obj = objects.NewBool(a >= b)
	default:
		log.Warn().Msgf("unexpected operator %d", operator)
	}

	return obj
}
