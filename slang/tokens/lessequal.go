package tokens

import "github.com/jamestunnell/go-synth/slang"

type LessEqual struct{}

const TypeLESSEQUAL = "LESSEQUAL"

func LESSEQUAL() slang.Token       { return &LessEqual{} }
func (t *LessEqual) Type() string  { return TypeLESSEQUAL }
func (t *LessEqual) Value() string { return "<=" }
