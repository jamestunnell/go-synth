package tokens

import "github.com/jamestunnell/go-synth/slang"

type LParen struct{}

const StrLPAREN = "("

func LPAREN() slang.TokenInfo           { return &LParen{} }
func (t *LParen) Type() slang.TokenType { return slang.TokenLPAREN }
func (t *LParen) Value() string         { return StrLPAREN }
