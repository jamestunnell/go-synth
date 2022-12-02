package tokens

import "github.com/jamestunnell/go-synth/slang"

type False struct{}

const StrFALSE = "false"

func FALSE() slang.Token       { return &False{} }
func (t *False) Type() string  { return "FALSE" }
func (t *False) Value() string { return StrFALSE }
