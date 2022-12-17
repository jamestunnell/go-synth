package interpreter_test

import (
	"strings"
	"testing"

	"github.com/jamestunnell/go-synth/slang"
	"github.com/jamestunnell/go-synth/slang/interpreter"
	"github.com/jamestunnell/go-synth/slang/lexer"
	"github.com/jamestunnell/go-synth/slang/objects"
	"github.com/jamestunnell/go-synth/slang/parser"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEvalExpression(t *testing.T) {
	// plain ints
	testExpr(t, "5", objects.NewInteger(5))
	testExpr(t, "12003", objects.NewInteger(12003))
	testExpr(t, "-54321", objects.NewInteger(-54321))

	// plain bools
	testExpr(t, "true", objects.NewBool(true))
	testExpr(t, "false", objects.NewBool(false))

	// integer arithmetic
	testExpr(t, "12 + 11", objects.NewInteger(23))
	testExpr(t, "1200 - 200", objects.NewInteger(1000))
	testExpr(t, "50 * 22", objects.NewInteger(1100))
	testExpr(t, "140 / 14", objects.NewInteger(10))
	testExpr(t, "2 + 12 * 10", objects.NewInteger(122))
	testExpr(t, "(2 + 12) * 10", objects.NewInteger(140))
	testExpr(t, "-(100 + 10)", objects.NewInteger(-110))
	testExpr(t, "10 * 10 * 10 * 10", objects.NewInteger(10000))

	// integer comparisons
	testExpr(t, "100 == 99", objects.NewBool(false))
	testExpr(t, "100 == 100", objects.NewBool(true))
	testExpr(t, "100 >= 99", objects.NewBool(true))
	testExpr(t, "100 >= 100", objects.NewBool(true))
	testExpr(t, "100 >= 101", objects.NewBool(false))
	testExpr(t, "100 > 99", objects.NewBool(true))
	testExpr(t, "100 > 100", objects.NewBool(false))
	testExpr(t, "100 < 101", objects.NewBool(true))
	testExpr(t, "100 < 100", objects.NewBool(false))
	testExpr(t, "100 < 99", objects.NewBool(false))
	testExpr(t, "100 <= 101", objects.NewBool(true))
	testExpr(t, "100 <= 100", objects.NewBool(true))
	testExpr(t, "100 <= 99", objects.NewBool(false))
}

func testExpr(t *testing.T, input string, expected slang.Object) {
	t.Run(input, func(t *testing.T) {
		l := lexer.New(strings.NewReader(input))
		p := parser.New(l)
		results := p.Run()

		for _, pErr := range results.Errors {
			t.Logf("parse error at %s: %v", pErr.Token.Location, pErr.Error)
		}

		require.Empty(t, results.Errors)
		require.Len(t, results.Statements, 1)

		obj := interpreter.EvalStatement(results.Statements[0])

		assert.Equal(t, expected.Inspect(), obj.Inspect())
	})
}
