package tokens

import "github.com/jamestunnell/go-synth/slang"

type Let struct{}

const (
	StrLET = "let"

	TypeLET = "LET"
)

func LET() slang.Token       { return &Let{} }
func (t *Let) Type() string  { return TypeLET }
func (t *Let) Value() string { return StrLET }
