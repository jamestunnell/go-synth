package parser_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/jamestunnell/go-synth/slang/lexer"
	"github.com/jamestunnell/go-synth/slang/parser"
)

func TestParser(t *testing.T) {
	input := `
let x  = 5
let y = 10.55 - 54
let foobar = func() {
	return 4
}
return (50 - 2)
`
	r := strings.NewReader(input)
	l := lexer.New(r)
	p := parser.New(l)

	prog, err := p.ParseProgram()

	require.NoError(t, err)
	assert.Len(t, prog.Statements, 4)
}
