package tokens

import "github.com/jamestunnell/go-synth/slang"

type Not struct{}

func NOT() slang.TokenInfo           { return &Not{} }
func (t *Not) Type() slang.TokenType { return slang.TokenNOT }
func (t *Not) Value() string         { return "!" }
