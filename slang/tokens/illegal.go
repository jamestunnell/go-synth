package tokens

import "github.com/jamestunnell/go-synth/slang"

type Illegal struct{ val rune }

func ILLEGAL(val rune) slang.Token       { return &Illegal{val: val} }
func (t *Illegal) Type() slang.TokenType { return slang.TokenILLEGAL }
func (t *Illegal) Value() string         { return string([]rune{t.val}) }
