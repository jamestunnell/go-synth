package tokens

import "github.com/jamestunnell/go-synth/slang"

type Illegal struct{ val rune }

func ILLEGAL(val rune) slang.Token { return &Illegal{val: val} }
func (t *Illegal) Type() string    { return "ILLEGAL" }
func (t *Illegal) String() string  { return string([]rune{t.val}) }
