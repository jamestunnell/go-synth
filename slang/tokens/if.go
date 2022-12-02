package tokens

import "github.com/jamestunnell/go-synth/slang"

type If struct{}

const (
	StrIF = "if"

	TypeIF = "IF"
)

func IF() slang.Token       { return &If{} }
func (t *If) Type() string  { return TypeIF }
func (t *If) Value() string { return StrIF }
