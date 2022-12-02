package tokens

import "github.com/jamestunnell/go-synth/slang"

type LBrace struct{}

const TypeLBRACE = "LBRACE"

func LBRACE() slang.Token       { return &LBrace{} }
func (t *LBrace) Type() string  { return TypeLBRACE }
func (t *LBrace) Value() string { return "{" }
