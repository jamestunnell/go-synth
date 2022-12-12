package tokens

import "github.com/jamestunnell/go-synth/slang"

type LessEqual struct{}

func LESSEQUAL() slang.Token               { return &LessEqual{} }
func (t *LessEqual) Type() slang.TokenType { return slang.TokenLESSEQUAL }
func (t *LessEqual) Value() string         { return "<=" }
