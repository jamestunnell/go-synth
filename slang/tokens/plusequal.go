package tokens

import "github.com/jamestunnell/go-synth/slang"

type PlusEqual struct{}

func PLUSEQUAL() slang.Token       { return &PlusEqual{} }
func (t *PlusEqual) Type() string  { return "PLUSEQUAL" }
func (t *PlusEqual) Value() string { return "+=" }
