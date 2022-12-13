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

func TestParserAssignMissingValue(t *testing.T) {
	testParserErrs(t, "x = ")
}

func TestParserAssignMissingEqual(t *testing.T) {
	testParserErrs(t, "x 5")
}

func TestParserGroupedExpr(t *testing.T) {
	testParser(t, "x = (4 + y)")
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

func TestParserIfStatement(t *testing.T) {
	const input = `if a == 2 {
		x = 10
	}`
	cond := expressions.NewEqual(
		expressions.NewIdentifier("a"),
		expressions.NewInteger(2))
	assign := statements.NewAssign(
		expressions.NewIdentifier("x"),
		expressions.NewInteger(10))
	ifStmnt := statements.NewIf(cond, assign)

	testParser(t, input, ifStmnt)

}

func testParser(t *testing.T, input string, expected ...slang.Statement) {
	results := parser.Parse(input)

	if !assert.Empty(t, results.Errors) {
		for i, err := range results.Errors {
			log.Debug().
				Err(err.Error).
				Int("line", err.Token.Location.Line).
				Int("column", err.Token.Location.Column).
				Str("context", err.Context.Description).
				Str("token", err.Token.Info.Value()).
				Msgf("parse error #%d", i+1)
		}
	}
	assert.Len(t, results.Statements, len(expected))

	for i := 0; i < len(results.Statements); i++ {
		s := results.Statements[i]

		assert.Equal(t, expected[i].Type(), s.Type())
		assert.True(t, s.Equal(expected[i]))
	}
}

func testParserErrs(t *testing.T, input string) {
	results := parser.Parse(input)

	require.NotEmpty(t, results.Errors)
}
