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
	testLexer(t, "var_1", slang.NewToken(tokens.IDENT("var_1"), slang.NewLoc(1, 1)))
}

func TestLexer_AssignInt(t *testing.T) {
	expected := []*slang.Token{
		slang.NewToken(tokens.IDENT("x"), slang.NewLoc(1, 4)),
		slang.NewToken(tokens.ASSIGN(), slang.NewLoc(1, 5)),
		slang.NewToken(tokens.INT("5"), slang.NewLoc(1, 6)),
	}

	testLexer(t, "   x=5", expected...)
	// testLexer(t, "\t \tx   = 5   ", expected...)
}

// func TestLexer_StatementsWithNewline(t *testing.T) {
// 	const str = "x = 5\ny = 10"

// 	expected := []slang.Token{
// 		tokens.IDENT("x"), tokens.ASSIGN(), tokens.INT("5"), tokens.LINE(),
// 		tokens.IDENT("y"), tokens.ASSIGN(), tokens.INT("10"),
// 	}

// 	testLexer(t, str, expected...)
// }

// func TestLexer_FloatMath(t *testing.T) {
// 	expected := []slang.Token{
// 		tokens.IDENT("my_num"), tokens.ASSIGN(), tokens.LPAREN(),
// 		tokens.FLOAT("2.5"), tokens.PLUS(), tokens.FLOAT("7.7"),
// 		tokens.RPAREN(), tokens.STAR(), tokens.LPAREN(),
// 		tokens.IDENT("otherNum"), tokens.SLASH(), tokens.FLOAT("33.5"),
// 		tokens.RPAREN(),
// 	}

// 	testLexer(t, "my_num = (2.5 + 7.7) * (otherNum / 33.5)", expected...)
// }

// func TestLexer_AssignFunc(t *testing.T) {
// 	input := "y = func() { \n\treturn 7\n}"
// 	expected := []slang.Token{
// 		tokens.IDENT("y"), tokens.ASSIGN(), tokens.FUNC(),
// 		tokens.LPAREN(), tokens.RPAREN(), tokens.LBRACE(), tokens.LINE(),
// 		tokens.RETURN(), tokens.INT("7"), tokens.LINE(), tokens.RBRACE(),
// 	}

// 	testLexer(t, input, expected...)
// }

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
