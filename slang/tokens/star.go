package tokens

import "github.com/jamestunnell/go-synth/slang"

type Star struct{}

const (
	StrSTAR  = "*"
	TypeSTAR = "STAR"
)

func STAR() slang.Token       { return &Star{} }
func (t *Star) Type() string  { return TypeSTAR }
func (t *Star) Value() string { return StrSTAR }
