package tokens

import "github.com/jamestunnell/go-synth/slang"

type Greater struct{}

const TypeGREATER = "GREATER"

func GREATER() slang.Token       { return &Greater{} }
func (t *Greater) Type() string  { return TypeGREATER }
func (t *Greater) Value() string { return ">" }
