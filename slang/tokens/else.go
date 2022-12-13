package tokens

import "github.com/jamestunnell/go-synth/slang"

type Else struct{}

const StrELSE = "else"

func ELSE() slang.TokenInfo           { return &Else{} }
func (t *Else) Type() slang.TokenType { return slang.TokenELSE }
func (t *Else) Value() string         { return StrELSE }
