package tokens

import "github.com/jamestunnell/go-synth/slang"

type Eof struct{}

const TypeEOF = "EOF"

func EOF() slang.Token       { return &Eof{} }
func (t *Eof) Type() string  { return TypeEOF }
func (t *Eof) Value() string { return "" }
