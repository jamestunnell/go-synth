package tokens

import "github.com/jamestunnell/go-synth/slang"

type Less struct{}

func LESS() slang.Token       { return &Less{} }
func (t *Less) Type() string  { return "LESS" }
func (t *Less) Value() string { return "<" }
