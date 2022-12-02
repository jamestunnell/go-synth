package tokens

import "github.com/jamestunnell/go-synth/slang"

type Slash struct{}

const (
	StrSLASH  = "/"
	TypeSLASH = "SLASH"
)

func SLASH() slang.Token       { return &Slash{} }
func (t *Slash) Type() string  { return TypeSLASH }
func (t *Slash) Value() string { return StrSLASH }
