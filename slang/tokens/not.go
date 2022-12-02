package tokens

import "github.com/jamestunnell/go-synth/slang"

type Not struct{}

func NOT() slang.Token       { return &Not{} }
func (t *Not) Type() string  { return "NOT" }
func (t *Not) Value() string { return "!" }
