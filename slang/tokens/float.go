package tokens

import "github.com/jamestunnell/go-synth/slang"

type Float struct {
	val string
}

func FLOAT(val string) slang.TokenInfo { return &Float{val: val} }
func (t *Float) Type() slang.TokenType { return slang.TokenFLOAT }
func (t *Float) Value() string         { return t.val }
