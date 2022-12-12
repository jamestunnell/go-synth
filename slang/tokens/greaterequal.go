package tokens

import "github.com/jamestunnell/go-synth/slang"

type GreaterEqual struct{}

const TypeGREATEREQUAL = "GREATEREQUAL"

func GREATEREQUAL() slang.Token       { return &GreaterEqual{} }
func (t *GreaterEqual) Type() string  { return TypeGREATEREQUAL }
func (t *GreaterEqual) Value() string { return ">=" }
