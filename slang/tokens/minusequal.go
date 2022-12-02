package tokens

import "github.com/jamestunnell/go-synth/slang"

type MinusEqual struct{}

func MINUSEQUAL() slang.Token       { return &MinusEqual{} }
func (t *MinusEqual) Type() string  { return "MINUSEQUAL" }
func (t *MinusEqual) Value() string { return "-=" }
