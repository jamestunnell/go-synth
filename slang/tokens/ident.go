package tokens

import "github.com/jamestunnell/go-synth/slang"

type Ident struct {
	val string
}

func IDENT(val string) slang.Token { return &Ident{val: val} }
func (t *Ident) Type() string      { return "IDENT" }
func (t *Ident) String() string    { return t.val }
