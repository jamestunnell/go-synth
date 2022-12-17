package interpreter

import (
	"github.com/jamestunnell/go-synth/slang"
	"github.com/jamestunnell/go-synth/slang/expressions"
	"github.com/jamestunnell/go-synth/slang/objects"
	"github.com/jamestunnell/go-synth/slang/statements"
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
		return evalAdd(ee)
	case *expressions.Subtract:
		return evalSub(ee)
	case *expressions.Multiply:
		return evalMul(ee)
	case *expressions.Divide:
		return evalDiv(ee)
	case *expressions.Equal:
		return evalEqual(ee)
	case *expressions.GreaterEqual:
		return evalGreaterEqual(ee)
	case *expressions.LessEqual:
		return evalLessEqual(ee)
	case *expressions.Less:
		return evalLess(ee)
	case *expressions.Greater:
		return evalGreater(ee)
	case *expressions.NotEqual:
		return evalNotEqual(ee)
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
	}

	return NULL
}

func evalAdd(e *expressions.Add) slang.Object {
	l := EvalExpression(e.Left)
	r := EvalExpression(e.Right)

	if l.Type() == slang.ObjectINTEGER && r.Type() == slang.ObjectINTEGER {
		return objects.NewInteger(l.(*objects.Integer).Value + r.(*objects.Integer).Value)
	}

	return NULL
}

func evalSub(e *expressions.Subtract) slang.Object {
	l := EvalExpression(e.Left)
	r := EvalExpression(e.Right)

	if l.Type() == slang.ObjectINTEGER && r.Type() == slang.ObjectINTEGER {
		return objects.NewInteger(l.(*objects.Integer).Value - r.(*objects.Integer).Value)
	}

	return NULL
}

func evalMul(e *expressions.Multiply) slang.Object {
	l := EvalExpression(e.Left)
	r := EvalExpression(e.Right)

	if l.Type() == slang.ObjectINTEGER && r.Type() == slang.ObjectINTEGER {
		return objects.NewInteger(l.(*objects.Integer).Value * r.(*objects.Integer).Value)
	}

	return NULL
}

func evalDiv(e *expressions.Divide) slang.Object {
	l := EvalExpression(e.Left)
	r := EvalExpression(e.Right)

	if l.Type() == slang.ObjectINTEGER && r.Type() == slang.ObjectINTEGER {
		return objects.NewInteger(l.(*objects.Integer).Value / r.(*objects.Integer).Value)
	}

	return NULL
}

func evalEqual(e *expressions.Equal) slang.Object {
	l := EvalExpression(e.Left)
	r := EvalExpression(e.Right)

	if l.Type() == slang.ObjectINTEGER && r.Type() == slang.ObjectINTEGER {
		return objects.NewBool(l.(*objects.Integer).Value == r.(*objects.Integer).Value)
	}

	return NULL
}

func evalNotEqual(e *expressions.NotEqual) slang.Object {
	l := EvalExpression(e.Left)
	r := EvalExpression(e.Right)

	if l.Type() == slang.ObjectINTEGER && r.Type() == slang.ObjectINTEGER {
		return objects.NewBool(l.(*objects.Integer).Value != r.(*objects.Integer).Value)
	}

	return NULL
}

func evalLess(e *expressions.Less) slang.Object {
	l := EvalExpression(e.Left)
	r := EvalExpression(e.Right)

	if l.Type() == slang.ObjectINTEGER && r.Type() == slang.ObjectINTEGER {
		return objects.NewBool(l.(*objects.Integer).Value < r.(*objects.Integer).Value)
	}

	return NULL
}

func evalGreater(e *expressions.Greater) slang.Object {
	l := EvalExpression(e.Left)
	r := EvalExpression(e.Right)

	if l.Type() == slang.ObjectINTEGER && r.Type() == slang.ObjectINTEGER {
		return objects.NewBool(l.(*objects.Integer).Value > r.(*objects.Integer).Value)
	}

	return NULL
}

func evalGreaterEqual(e *expressions.GreaterEqual) slang.Object {
	l := EvalExpression(e.Left)
	r := EvalExpression(e.Right)

	if l.Type() == slang.ObjectINTEGER && r.Type() == slang.ObjectINTEGER {
		return objects.NewBool(l.(*objects.Integer).Value >= r.(*objects.Integer).Value)
	}

	return NULL
}

func evalLessEqual(e *expressions.LessEqual) slang.Object {
	l := EvalExpression(e.Left)
	r := EvalExpression(e.Right)

	if l.Type() == slang.ObjectINTEGER && r.Type() == slang.ObjectINTEGER {
		return objects.NewBool(l.(*objects.Integer).Value <= r.(*objects.Integer).Value)
	}

	return NULL
}
