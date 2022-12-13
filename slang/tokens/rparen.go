package tokens

import "github.com/jamestunnell/go-synth/slang"

type RParen struct{}

func RPAREN() slang.TokenInfo           { return &RParen{} }
func (t *RParen) Type() slang.TokenType { return slang.TokenRPAREN }
func (t *RParen) Value() string         { return ")" }
