package tokens

import "github.com/jamestunnell/go-synth/slang"

type Bang struct{}

func BANG() slang.TokenInfo           { return &Bang{} }
func (t *Bang) Type() slang.TokenType { return slang.TokenBANG }
func (t *Bang) Value() string         { return "!" }
