package tokens

import "github.com/jamestunnell/go-synth/slang"

type Minus struct{}

const StrMINUS = "-"

func MINUS() slang.TokenInfo           { return &Minus{} }
func (t *Minus) Type() slang.TokenType { return slang.TokenMINUS }
func (t *Minus) Value() string         { return StrMINUS }
