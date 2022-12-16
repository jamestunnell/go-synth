package lexer_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/jamestunnell/go-synth/slang"
	"github.com/jamestunnell/go-synth/slang/lexer"
	"github.com/jamestunnell/go-synth/slang/tokens"
)

func TestLexer_IndentWithDigits(t *testing.T) {
	testLexer(t, "var_1", tok(tokens.IDENT("var_1"), 1, 1))
}

func tok(info slang.TokenInfo, line, col int) *slang.Token {
	return slang.NewToken(info, slang.NewLoc(line, col))
}

func TestLexer_AssignInt(t *testing.T) {
	expected := []*slang.Token{
		tok(tokens.IDENT("x"), 1, 4),
		tok(tokens.ASSIGN(), 1, 5),
		tok(tokens.INT("5"), 1, 6),
	}

	testLexer(t, "   x=5", expected...)
	// testLexer(t, "\t \tx   = 5   ", expected...)
}

func TestLexer_StatementsWithNewline(t *testing.T) {
	const str = "x = 5\ny = 10"

	expected := []*slang.Token{
		tok(tokens.IDENT("x"), 1, 1),
		tok(tokens.ASSIGN(), 1, 3),
		tok(tokens.INT("5"), 1, 5),
		tok(tokens.LINE(), 1, 6),
		tok(tokens.IDENT("y"), 2, 1),
		tok(tokens.ASSIGN(), 2, 3),
		tok(tokens.INT("10"), 2, 5),
	}

	testLexer(t, str, expected...)
}

func TestLexer_FloatMath(t *testing.T) {
	expected := []*slang.Token{
		tok(tokens.IDENT("my_num"), 1, 1),
		tok(tokens.ASSIGN(), 1, 8),
		tok(tokens.LPAREN(), 1, 10),
		tok(tokens.FLOAT("2.5"), 1, 11),
		tok(tokens.PLUS(), 1, 15),
		tok(tokens.FLOAT("7.7"), 1, 17),
		tok(tokens.RPAREN(), 1, 20),
		tok(tokens.STAR(), 1, 22),
		tok(tokens.LPAREN(), 1, 24),
		tok(tokens.IDENT("otherNum"), 1, 25),
		tok(tokens.SLASH(), 1, 34),
		tok(tokens.FLOAT("33.5"), 1, 36),
		tok(tokens.RPAREN(), 1, 40),
	}

	testLexer(t, "my_num = (2.5 + 7.7) * (otherNum / 33.5)", expected...)
}

func TestLexer_AssignFunc(t *testing.T) {
	input := "y = func() { \n\treturn 7\n}"
	expected := []*slang.Token{
		tok(tokens.IDENT("y"), 1, 1),
		tok(tokens.ASSIGN(), 1, 3),
		tok(tokens.FUNC(), 1, 5),
		tok(tokens.LPAREN(), 1, 9),
		tok(tokens.RPAREN(), 1, 10),
		tok(tokens.LBRACE(), 1, 12),
		tok(tokens.LINE(), 1, 14),
		tok(tokens.RETURN(), 2, 2),
		tok(tokens.INT("7"), 2, 9),
		tok(tokens.LINE(), 2, 10),
		tok(tokens.RBRACE(), 3, 1),
	}

	testLexer(t, input, expected...)
}

func testLexer(t *testing.T, input string, expected ...*slang.Token) {
	toks := lexer.ScanString(input)

	require.Len(t, toks, len(expected))

	for i := 0; i < len(toks); i++ {
		testTokensEqual(t, expected[i], toks[i])
	}
}

func testTokensEqual(t *testing.T, exp, act *slang.Token) {
	assert.Equal(t, exp.Info.Type(), act.Info.Type())
	assert.Equal(t, exp.Info.Value(), act.Info.Value())
	assert.Equal(t, exp.Location.Line, act.Location.Line)
	assert.Equal(t, exp.Location.Column, act.Location.Column)
}
