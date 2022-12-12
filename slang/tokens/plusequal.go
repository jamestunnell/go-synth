package tokens

import "github.com/jamestunnell/go-synth/slang"

type PlusEqual struct{}

func PLUSEQUAL() slang.Token               { return &PlusEqual{} }
func (t *PlusEqual) Type() slang.TokenType { return slang.TokenPLUSEQUAL }
func (t *PlusEqual) Value() string         { return "+=" }
