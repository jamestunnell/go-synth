package tokens

import "github.com/jamestunnell/go-synth/slang"

type Assign struct{}

func ASSIGN() slang.Token        { return &Assign{} }
func (t *Assign) Type() string   { return "ASSIGN" }
func (t *Assign) String() string { return "=" }
