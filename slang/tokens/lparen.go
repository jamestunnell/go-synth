package tokens

import "github.com/jamestunnell/go-synth/slang"

type LParen struct{}

func LPAREN() slang.Token        { return &LParen{} }
func (t *LParen) Type() string   { return "LPAREN" }
func (t *LParen) String() string { return "(" }
