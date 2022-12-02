package tokens

import "github.com/jamestunnell/go-synth/slang"

type LessEqual struct{}

func LESSEQUAL() slang.Token       { return &LessEqual{} }
func (t *LessEqual) Type() string  { return "LESSEQUAL" }
func (t *LessEqual) Value() string { return "<=" }
