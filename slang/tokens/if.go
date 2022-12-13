package tokens

import "github.com/jamestunnell/go-synth/slang"

type If struct{}

const (
	StrIF = "if"
)

func IF() slang.TokenInfo           { return &If{} }
func (t *If) Type() slang.TokenType { return slang.TokenIF }
func (t *If) Value() string         { return StrIF }
