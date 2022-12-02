package tokens

import "github.com/jamestunnell/go-synth/slang"

type MinusMinus struct{}

func MINUSMINUS() slang.Token       { return &MinusMinus{} }
func (t *MinusMinus) Type() string  { return "MINUSMINUS" }
func (t *MinusMinus) Value() string { return "--" }
