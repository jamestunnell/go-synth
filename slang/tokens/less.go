package tokens

import "github.com/jamestunnell/go-synth/slang"

type Less struct{}

const TypeLESS = "LESS"

func LESS() slang.Token       { return &Less{} }
func (t *Less) Type() string  { return TypeLESS }
func (t *Less) Value() string { return "<" }
