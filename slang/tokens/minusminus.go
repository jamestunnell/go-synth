package tokens

import "github.com/jamestunnell/go-synth/slang"

type MinusMinus struct{}

func MINUSMINUS() slang.TokenInfo           { return &MinusMinus{} }
func (t *MinusMinus) Type() slang.TokenType { return slang.TokenMINUSMINUS }
func (t *MinusMinus) Value() string         { return "--" }
