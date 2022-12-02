package tokens

import "github.com/jamestunnell/go-synth/slang"

type RBrace struct{}

const TypeRBRACE = "RBRACE"

func RBRACE() slang.Token       { return &RBrace{} }
func (t *RBrace) Type() string  { return TypeRBRACE }
func (t *RBrace) Value() string { return "}" }
