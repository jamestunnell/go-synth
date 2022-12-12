package tokens

import "github.com/jamestunnell/go-synth/slang"

type Line struct{}

const TypeLINE = "LINE"

func LINE() slang.Token       { return &Line{} }
func (l *Line) Type() string  { return TypeLINE }
func (l *Line) Value() string { return "\n" }
