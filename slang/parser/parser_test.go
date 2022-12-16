package parser_test

import (
	"testing"

	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/jamestunnell/go-synth/slang"
	"github.com/jamestunnell/go-synth/slang/expressions"
	"github.com/jamestunnell/go-synth/slang/parser"
	"github.com/jamestunnell/go-synth/slang/statements"
)

func TestParserExprStatement(t *testing.T) {
	testCases := map[string]slang.Statement{
		// plain values
		"x":     se(id("x")),
		"5":     se(expressions.NewInteger(5)),
		"25.7":  se(expressions.NewFloat(25.7)),
		"false": se(expressions.NewBool(false)),
		"true":  se(expressions.NewBool(true)),

		// prefix operators
		"-15":   se(expressions.NewNegative(expressions.NewInteger(15))),
		"!true": se(expressions.NewNot(expressions.NewBool(true))),

		// infix operators
		"a + b":  se(add(id("a"), id("b"))),
		"a - b":  se(sub(id("a"), id("b"))),
		"a * b":  se(mul(id("a"), id("b"))),
		"a / b":  se(div(id("a"), id("b"))),
		"a > b":  se(expressions.NewGreater(id("a"), id("b"))),
		"a < b":  se(expressions.NewLess(id("a"), id("b"))),
		"a == b": se(expressions.NewEqual(id("a"), id("b"))),
		"a != b": se(expressions.NewNotEqual(id("a"), id("b"))),

		// more ellaborate expressions
		"10 + 7 - 3":    se(sub(add(i(10), i(7)), i(3))),
		"15 + 2 * 12":   se(add(i(15), mul(i(2), i(12)))),
		"6 * 6 - 3 * 3": se(sub(mul(i(6), i(6)), mul(i(3), i(3)))),

		// grouped expression
		"(15 + 2) * 12": se(mul(add(i(15), i(2)), i(12))),

		// func calls
		"sum(1,2,3)":     se(call(id("sum"), i(1), i(2), i(3))),
		"5 * sub(10, 5)": se(mul(i(5), call(id("sub"), i(10), i(5)))),
	}

	for input, expected := range testCases {
		t.Run(input, func(t *testing.T) {
			testParser(t, input, expected)
		})
	}
}

func TestParserAssignMissingValue(t *testing.T) {
	testParserErrs(t, "x = ")
}

func TestParserAssignMissingEqual(t *testing.T) {
	testParserErrs(t, "x 5")
}

func TestParserGroupedExpr(t *testing.T) {
	s := statements.NewAssign(
		expressions.NewIdentifier("x"),
		expressions.NewAdd(
			expressions.NewInteger(4),
			expressions.NewIdentifier("y"),
		),
	)

	testParser(t, "x = (4 + y)", s)
}

func TestParserOneAssignStatement(t *testing.T) {
	identExpr := expressions.NewIdentifier("x")
	intExpr := expressions.NewInteger(5)
	assign := statements.NewAssign(identExpr, intExpr)

	testParser(t, "x = 5", assign)
}

func TestParserThreeAssignStatements(t *testing.T) {
	const input = `a = 77
	b = 100.0
	longer_name = 75.0 - 22.2`

	a := expressions.NewIdentifier("a")
	b := expressions.NewIdentifier("b")
	c := expressions.NewIdentifier("longer_name")

	aVal := expressions.NewInteger(77)
	bVal := expressions.NewFloat(100.0)
	cVal := expressions.NewSubtract(
		expressions.NewFloat(75.0),
		expressions.NewFloat(22.2))

	testParser(t, input,
		statements.NewAssign(a, aVal),
		statements.NewAssign(b, bVal),
		statements.NewAssign(c, cVal))
}

func TestParserReturnStatement(t *testing.T) {
	l := expressions.NewFloat(12.77)
	r := expressions.NewIdentifier("num")
	add := expressions.NewAdd(l, r)
	ret := statements.NewReturn(add)

	testParser(t, "return 12.77 + num", ret)
}

func TestParserIfExpr(t *testing.T) {
	input := `y = if a == 2 {
		x + 10
	}`
	cond := expressions.NewEqual(
		expressions.NewIdentifier("a"),
		expressions.NewInteger(2))
	assign := statements.NewExpression(
		expressions.NewAdd(
			expressions.NewIdentifier("x"),
			expressions.NewInteger(10)))
	conseq := statements.NewBlock(assign)
	ifExpr := expressions.NewIf(cond, conseq)

	testParser(t, input, statements.NewAssign(id("y"), ifExpr))

	input += ` else {
		76
	}`

	altern := statements.NewBlock(
		statements.NewExpression(
			expressions.NewInteger(76),
		),
	)
	ifElseExpr := expressions.NewIfElse(cond, conseq, altern)

	testParser(t, input, statements.NewAssign(id("y"), ifElseExpr))
}

func TestParserFuncLiteralNoParams(t *testing.T) {
	const input = `myvar = func(){
		return 7
	}`

	body := statements.NewBlock(
		statements.NewReturn(expressions.NewInteger(7)),
	)
	af := expressions.NewFunctionLiteral(
		[]*expressions.Identifier{}, body)
	assign := statements.NewAssign(
		expressions.NewIdentifier("myvar"), af)

	testParser(t, input, assign)
}

func TestParserFuncLiteralOneParams(t *testing.T) {
	const input = `myvar = func(x){
		return 7
	}`

	body := statements.NewBlock(
		statements.NewReturn(expressions.NewInteger(7)),
	)
	af := expressions.NewFunctionLiteral(
		[]*expressions.Identifier{expressions.NewIdentifier("x")}, body)
	assign := statements.NewAssign(
		expressions.NewIdentifier("myvar"), af)

	testParser(t, input, assign)
}

func TestParserFuncLiteralTwoParams(t *testing.T) {
	const input = `myvar = func(x, y){
		return 7
	}`

	body := statements.NewBlock(
		statements.NewReturn(expressions.NewInteger(7)),
	)
	af := expressions.NewFunctionLiteral(
		[]*expressions.Identifier{
			expressions.NewIdentifier("x"),
			expressions.NewIdentifier("y"),
		}, body)
	assign := statements.NewAssign(
		expressions.NewIdentifier("myvar"), af)

	testParser(t, input, assign)
}

func testParser(t *testing.T, input string, expected ...slang.Statement) {
	results := parser.Parse(input)

	for i, err := range results.Errors {
		log.Debug().
			Err(err.Error).
			Int("line", err.Token.Location.Line).
			Int("column", err.Token.Location.Column).
			Str("context", err.Context.Description).
			Str("token", err.Token.Info.Value()).
			Msgf("parse error #%d", i+1)
	}

	require.Empty(t, results.Errors)

	require.Len(t, results.Statements, len(expected))

	for i := 0; i < len(results.Statements); i++ {
		s := results.Statements[i]

		assert.Equal(t, expected[i].Type(), s.Type())
		if !assert.True(t, s.Equal(expected[i])) {
			t.Logf("statements not equal: expected %#v, got %#v", expected[i], s)
		}
	}
}

func testParserErrs(t *testing.T, input string) {
	results := parser.Parse(input)

	require.NotEmpty(t, results.Errors)

	for _, pErr := range results.Errors {
		t.Logf("parse error at %s: %v", pErr.Token.Location, pErr.Error)
	}
}

func se(expr slang.Expression) slang.Statement {
	return statements.NewExpression(expr)
}

func id(name string) slang.Expression {
	return expressions.NewIdentifier(name)
}

func add(left, right slang.Expression) slang.Expression {
	return expressions.NewAdd(left, right)
}

func sub(left, right slang.Expression) slang.Expression {
	return expressions.NewSubtract(left, right)
}

func mul(left, right slang.Expression) slang.Expression {
	return expressions.NewMultiply(left, right)
}

func div(left, right slang.Expression) slang.Expression {
	return expressions.NewDivide(left, right)
}

func i(val int64) slang.Expression {
	return expressions.NewInteger(val)
}

func call(fn slang.Expression, args ...slang.Expression) slang.Expression {
	return expressions.NewCall(fn, args...)
}
