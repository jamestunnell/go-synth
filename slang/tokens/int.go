package tokens

import "github.com/jamestunnell/go-synth/slang"

type Int struct {
	val string
}

func INT(val string) slang.TokenInfo { return &Int{val: val} }
func (t *Int) Type() slang.TokenType { return slang.TokenINT }
func (t *Int) Value() string         { return t.val }
