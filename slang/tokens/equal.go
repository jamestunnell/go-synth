package tokens

import "github.com/jamestunnell/go-synth/slang"

type Equal struct{}

const TypeEQUAL = "EQUAL"

func EQUAL() slang.Token       { return &Equal{} }
func (t *Equal) Type() string  { return TypeEQUAL }
func (t *Equal) Value() string { return "==" }
