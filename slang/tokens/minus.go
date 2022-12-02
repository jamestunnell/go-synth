package tokens

import "github.com/jamestunnell/go-synth/slang"

type Minus struct{}

const (
	StrMINUS  = "-"
	TypeMINUS = "MINUS"
)

func MINUS() slang.Token       { return &Minus{} }
func (t *Minus) Type() string  { return TypeMINUS }
func (t *Minus) Value() string { return StrMINUS }
