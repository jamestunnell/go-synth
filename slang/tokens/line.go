package tokens

import "github.com/jamestunnell/go-synth/slang"

type Line struct{}

func LINE() slang.TokenInfo           { return &Line{} }
func (l *Line) Type() slang.TokenType { return slang.TokenLINE }
func (l *Line) Value() string         { return "\n" }
