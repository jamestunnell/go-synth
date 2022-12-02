package tokens

import "github.com/jamestunnell/go-synth/slang"

type Greater struct{}

func GREATER() slang.Token       { return &Greater{} }
func (t *Greater) Type() string  { return "GREATER" }
func (t *Greater) Value() string { return ">" }
