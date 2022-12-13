package tokens

import "github.com/jamestunnell/go-synth/slang"

type SlashEqual struct{}

const (
	StrSLASHEQUAL = "/="
)

func SLASHEQUAL() slang.TokenInfo           { return &SlashEqual{} }
func (t *SlashEqual) Type() slang.TokenType { return slang.TokenSLASHEQUAL }
func (t *SlashEqual) Value() string         { return StrSLASHEQUAL }
