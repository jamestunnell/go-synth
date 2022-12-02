package tokens

import "github.com/jamestunnell/go-synth/slang"

type GreaterEqual struct{}

func GREATEREQUAL() slang.Token       { return &GreaterEqual{} }
func (t *GreaterEqual) Type() string  { return "GREATEREQUAL" }
func (t *GreaterEqual) Value() string { return ">=" }
