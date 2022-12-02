package tokens

import "github.com/jamestunnell/go-synth/slang"

type PlusPlus struct{}

func PLUSPLUS() slang.Token       { return &PlusPlus{} }
func (t *PlusPlus) Type() string  { return "PLUSPLUS" }
func (t *PlusPlus) Value() string { return "++" }
