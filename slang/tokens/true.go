package tokens

import "github.com/jamestunnell/go-synth/slang"

type True struct{}

const StrTRUE = "true"

func TRUE() slang.Token       { return &True{} }
func (t *True) Type() string  { return "TRUE" }
func (t *True) Value() string { return StrTRUE }
