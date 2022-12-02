package tokens

import "github.com/jamestunnell/go-synth/slang"

type Assign struct{}

const (
	StrASSIGN = "="

	TypeASSIGN = "ASSIGN"
)

func ASSIGN() slang.Token       { return &Assign{} }
func (t *Assign) Type() string  { return TypeASSIGN }
func (t *Assign) Value() string { return StrASSIGN }
