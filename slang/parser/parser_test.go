package parser_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/jamestunnell/go-synth/slang"
	"github.com/jamestunnell/go-synth/slang/expressions"
	"github.com/jamestunnell/go-synth/slang/lexer"
	"github.com/jamestunnell/go-synth/slang/parser"
	"github.com/jamestunnell/go-synth/slang/statements"
)

func TestParserAssignStatement(t *testing.T) {
	identExpr := expressions.NewIdentifier("x")
	intExpr := expressions.NewInteger(5)
	assign := statements.NewAssign(identExpr, intExpr)

	testParser(t, "x = 5", assign)
}

func TestParserReturnStatement(t *testing.T) {
	l := expressions.NewFloat(12.77)
	r := expressions.NewIdentifier("num")
	add := expressions.NewBinaryOperation(expressions.Add, l, r)
	ret := statements.NewReturn(add)

	testParser(t, "return 12.77 + num", ret)

}

func testParser(t *testing.T, input string, expected ...slang.Statement) {
	r := strings.NewReader(input)
	l := lexer.New(r)
	p := parser.New(l)

	prog, err := p.ParseProgram()

	require.NoError(t, err)
	assert.Len(t, prog.Statements, len(expected))

	for i := 0; i < len(prog.Statements); i++ {
		s := prog.Statements[i]

		assert.Equal(t, expected[i].Type(), s.Type())
		assert.Equal(t, expected[i].String(), s.String())
	}
}
