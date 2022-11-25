package tokens

import "github.com/jamestunnell/go-synth/slang"

type Return struct{}

const StrRETURN = "return"

func RETURN() slang.Token        { return &Return{} }
func (t *Return) Type() string   { return "RETURN" }
func (t *Return) String() string { return StrRETURN }
