package tokens

import "github.com/jamestunnell/go-synth/slang"

type NotEqual struct{}

func NOTEQUAL() slang.Token       { return &NotEqual{} }
func (t *NotEqual) Type() string  { return "NOTEQUAL" }
func (t *NotEqual) Value() string { return "!=" }
