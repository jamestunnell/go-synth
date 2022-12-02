package tokens

import "github.com/jamestunnell/go-synth/slang"

type RParen struct{}

const TypeRPAREN = "RPAREn"

func RPAREN() slang.Token       { return &RParen{} }
func (t *RParen) Type() string  { return TypeRPAREN }
func (t *RParen) Value() string { return ")" }
