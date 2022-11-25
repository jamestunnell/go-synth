package tokens

import "github.com/jamestunnell/go-synth/slang"

type LBrace struct{}

func LBRACE() slang.Token        { return &LBrace{} }
func (t *LBrace) Type() string   { return "LBRACE" }
func (t *LBrace) String() string { return "{" }
