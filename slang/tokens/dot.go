package tokens

import "github.com/jamestunnell/go-synth/slang"

type Dot struct{}

func DOT() slang.TokenInfo           { return &Dot{} }
func (t *Dot) Type() slang.TokenType { return slang.TokenDOT }
func (t *Dot) Value() string         { return "." }
