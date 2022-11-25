package tokens

import "github.com/jamestunnell/go-synth/slang"

type RBrace struct{}

func RBRACE() slang.Token        { return &RBrace{} }
func (t *RBrace) Type() string   { return "RBRACE" }
func (t *RBrace) String() string { return "}" }
