package tokens

import "github.com/jamestunnell/go-synth/slang"

type Let struct{}

const StrLET = "let"

func LET() slang.Token        { return &Let{} }
func (t *Let) Type() string   { return "LET" }
func (t *Let) String() string { return StrLET }
