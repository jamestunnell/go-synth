package tokens

import "github.com/jamestunnell/go-synth/slang"

type Semicolon struct{}

func SEMICOLON() slang.Token        { return &Semicolon{} }
func (t *Semicolon) Type() string   { return "SEMICOLON" }
func (t *Semicolon) String() string { return ";" }
