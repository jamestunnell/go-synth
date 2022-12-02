package tokens

import "github.com/jamestunnell/go-synth/slang"

type Plus struct{}

const (
	StrPLUS  = "+"
	TypePLUS = "PLUS"
)

func PLUS() slang.Token       { return &Plus{} }
func (t *Plus) Type() string  { return TypePLUS }
func (t *Plus) Value() string { return StrPLUS }
