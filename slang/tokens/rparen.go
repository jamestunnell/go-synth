package tokens

import "github.com/jamestunnell/go-synth/slang"

type RParen struct{}

func RPAREN() slang.Token        { return &RParen{} }
func (t *RParen) Type() string   { return "RPAREN" }
func (t *RParen) String() string { return ")" }
