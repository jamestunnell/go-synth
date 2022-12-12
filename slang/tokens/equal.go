package tokens

import "github.com/jamestunnell/go-synth/slang"

type Equal struct{}

func EQUAL() slang.Token               { return &Equal{} }
func (t *Equal) Type() slang.TokenType { return slang.TokenEQUAL }
func (t *Equal) Value() string         { return "==" }
