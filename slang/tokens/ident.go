package tokens

import "github.com/jamestunnell/go-synth/slang"

type Ident struct {
	val string
}

const (
	TypeIDENT = "IDENT"
)

func IDENT(val string) slang.Token { return &Ident{val: val} }
func (t *Ident) Type() string      { return TypeIDENT }
func (t *Ident) Value() string     { return t.val }
