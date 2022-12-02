package tokens

import "github.com/jamestunnell/go-synth/slang"

type Func struct{}

const (
	StrFUNC  = "func"
	TypeFUNC = "FUNC"
)

func FUNC() slang.Token       { return &Func{} }
func (t *Func) Type() string  { return TypeFUNC }
func (t *Func) Value() string { return StrFUNC }
