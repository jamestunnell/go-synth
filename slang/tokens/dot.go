package tokens

import "github.com/jamestunnell/go-synth/slang"

type Dot struct{}

func DOT() slang.Token       { return &Dot{} }
func (t *Dot) Type() string  { return "DOT" }
func (t *Dot) Value() string { return "." }
