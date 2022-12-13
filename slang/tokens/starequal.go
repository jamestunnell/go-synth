package tokens

import "github.com/jamestunnell/go-synth/slang"

type StarEqual struct{}

const (
	StrSTAREQUAL = "*="
)

func STAREQUAL() slang.TokenInfo           { return &StarEqual{} }
func (t *StarEqual) Type() slang.TokenType { return slang.TokenSTAREQUAL }
func (t *StarEqual) Value() string         { return StrSTAREQUAL }
