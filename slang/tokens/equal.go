package tokens

import "github.com/jamestunnell/go-synth/slang"

type Equal struct{}

func EQUAL() slang.Token       { return &Equal{} }
func (t *Equal) Type() string  { return "EQUAL" }
func (t *Equal) Value() string { return "==" }
