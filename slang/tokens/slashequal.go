package tokens

import "github.com/jamestunnell/go-synth/slang"

type SlashEqual struct{}

const (
	StrSLASHEQUAL  = "/="
	TypeSLASHEQUAL = "SLASHEQUAL"
)

func SLASHEQUAL() slang.Token       { return &SlashEqual{} }
func (t *SlashEqual) Type() string  { return TypeSLASHEQUAL }
func (t *SlashEqual) Value() string { return StrSLASHEQUAL }
