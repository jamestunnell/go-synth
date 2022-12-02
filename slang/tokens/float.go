package tokens

import "github.com/jamestunnell/go-synth/slang"

type Float struct {
	val string
}

const TypeFLOAT = "FLOAT"

func FLOAT(val string) slang.Token { return &Float{val: val} }
func (t *Float) Type() string      { return TypeFLOAT }
func (t *Float) Value() string     { return t.val }
