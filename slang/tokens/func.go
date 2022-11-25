package tokens

import "github.com/jamestunnell/go-synth/slang"

type Func struct{}

const StrFUNC = "func"

func FUNC() slang.Token        { return &Func{} }
func (t *Func) Type() string   { return "FUNC" }
func (t *Func) String() string { return StrFUNC }
