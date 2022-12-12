package tokens

import "github.com/jamestunnell/go-synth/slang"

type GreaterEqual struct{}

func GREATEREQUAL() slang.Token               { return &GreaterEqual{} }
func (t *GreaterEqual) Type() slang.TokenType { return slang.TokenGREATEREQUAL }
func (t *GreaterEqual) Value() string         { return ">=" }
