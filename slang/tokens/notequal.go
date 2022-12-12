package tokens

import "github.com/jamestunnell/go-synth/slang"

type NotEqual struct{}

const TypeNOTEQUAL = "NOTEQUAL"

func NOTEQUAL() slang.Token       { return &NotEqual{} }
func (t *NotEqual) Type() string  { return TypeNOTEQUAL }
func (t *NotEqual) Value() string { return "!=" }
