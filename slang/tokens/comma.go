package tokens

import "github.com/jamestunnell/go-synth/slang"

type Comma struct{}

const TypeCOMMA = "COMMA"

func COMMA() slang.Token       { return &Comma{} }
func (t *Comma) Type() string  { return TypeCOMMA }
func (t *Comma) Value() string { return "," }
