package tokens

import "github.com/jamestunnell/go-synth/slang"

type Int struct {
	val string
}

func INT(val string) slang.Token { return &Int{val: val} }
func (t *Int) Type() string      { return "INT" }
func (t *Int) String() string    { return t.val }
