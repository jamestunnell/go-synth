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

func TestParserOneLine(t *testing.T) {
	testParser(t, "x = 5", statements.NewAssign(
		expressions.NewIdentifier("x"),
		expressions.NewInteger(5),
	))
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
