package tokens

import "github.com/jamestunnell/go-synth/slang"

type Float struct {
	val string
}

func FLOAT(val string) slang.Token { return &Float{val: val} }
func (t *Float) Type() string      { return "FLOAT" }
func (t *Float) String() string    { return t.val }
