package tokens

import "github.com/jamestunnell/go-synth/slang"

type Ident struct {
	val string
}

func IDENT(val string) slang.TokenInfo { return &Ident{val: val} }
func (t *Ident) Type() slang.TokenType { return slang.TokenIDENT }
func (t *Ident) Value() string         { return t.val }
