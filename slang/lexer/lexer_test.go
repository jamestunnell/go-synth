package lexer_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/jamestunnell/go-synth/slang"
	"github.com/jamestunnell/go-synth/slang/lexer"
	"github.com/jamestunnell/go-synth/slang/tokens"
)

func TestLexer_AssignInt(t *testing.T) {
	expected := []slang.Token{tokens.LET(), tokens.IDENT("x"),
		tokens.ASSIGN(), tokens.INT("5")}

	testLexer(t, "let x = 5", expected...)
	testLexer(t, " \n let x=5", expected...)
	testLexer(t, "\t let\tx   = 5   ", expected...)
}

func TestLexer_FloatMath(t *testing.T) {
	expected := []slang.Token{
		tokens.IDENT("my_num"), tokens.ASSIGN(), tokens.LPAREN(),
		tokens.FLOAT("2.5"), tokens.PLUS(), tokens.FLOAT("7.7"),
		tokens.RPAREN(), tokens.STAR(), tokens.LPAREN(),
		tokens.IDENT("otherNum"), tokens.SLASH(), tokens.FLOAT("33.5"),
		tokens.RPAREN(),
	}

	testLexer(t, "my_num = (2.5 + 7.7) * (otherNum / 33.5)", expected...)
}

func TestLexer_AssignFunc(t *testing.T) {
	input := "let y = func() { return 7 }"
	expected := []slang.Token{
		tokens.LET(), tokens.IDENT("y"), tokens.ASSIGN(), tokens.FUNC(),
		tokens.LPAREN(), tokens.RPAREN(), tokens.LBRACE(), tokens.RETURN(),
		tokens.INT("7"), tokens.RBRACE(),
	}

	testLexer(t, input, expected...)
}

func testLexer(t *testing.T, input string, expected ...slang.Token) {
	toks := lexer.ScanString(input)

	require.Len(t, toks, len(expected))

	for i := 0; i < len(toks); i++ {
		assert.Equal(t, expected[i].Type(), toks[i].Type())
		assert.Equal(t, expected[i].Value(), toks[i].Value())
	}
}
