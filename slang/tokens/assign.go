package tokens

import "github.com/jamestunnell/go-synth/slang"

type Assign struct{}

const StrASSIGN = "="

func ASSIGN() slang.Token               { return &Assign{} }
func (t *Assign) Type() slang.TokenType { return slang.TokenASSIGN }
func (t *Assign) Value() string         { return StrASSIGN }
