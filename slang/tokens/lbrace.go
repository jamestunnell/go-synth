package tokens

import "github.com/jamestunnell/go-synth/slang"

type LBrace struct{}

func LBRACE() slang.TokenInfo           { return &LBrace{} }
func (t *LBrace) Type() slang.TokenType { return slang.TokenLBRACE }
func (t *LBrace) Value() string         { return "{" }
