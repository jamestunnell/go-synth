package tokens

import "github.com/jamestunnell/go-synth/slang"

type Else struct{}

const StrELSE = "else"

func ELSE() slang.Token       { return &Else{} }
func (t *Else) Type() string  { return "ELSE" }
func (t *Else) Value() string { return StrELSE }
