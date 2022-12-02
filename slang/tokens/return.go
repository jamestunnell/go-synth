package tokens

import "github.com/jamestunnell/go-synth/slang"

type Return struct{}

const (
	StrRETURN = "return"

	TypeRETURN = "RETURN"
)

func RETURN() slang.Token       { return &Return{} }
func (t *Return) Type() string  { return TypeRETURN }
func (t *Return) Value() string { return StrRETURN }
