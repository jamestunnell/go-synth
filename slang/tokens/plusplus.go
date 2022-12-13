package tokens

import "github.com/jamestunnell/go-synth/slang"

type PlusPlus struct{}

func PLUSPLUS() slang.TokenInfo           { return &PlusPlus{} }
func (t *PlusPlus) Type() slang.TokenType { return slang.TokenPLUSPLUS }
func (t *PlusPlus) Value() string         { return "++" }
