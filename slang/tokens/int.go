package tokens

import "github.com/jamestunnell/go-synth/slang"

type Int struct {
	val string
}

const TypeINT = "INT"

func INT(val string) slang.Token { return &Int{val: val} }
func (t *Int) Type() string      { return TypeINT }
func (t *Int) Value() string     { return t.val }
