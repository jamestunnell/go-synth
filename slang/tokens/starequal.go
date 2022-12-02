package tokens

import "github.com/jamestunnell/go-synth/slang"

type StarEqual struct{}

const (
	StrSTAREQUAL  = "*="
	TypeSTAREQUAL = "STAREQUAL"
)

func STAREQUAL() slang.Token       { return &StarEqual{} }
func (t *StarEqual) Type() string  { return TypeSTAREQUAL }
func (t *StarEqual) Value() string { return StrSTAREQUAL }
