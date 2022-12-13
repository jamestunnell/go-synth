package tokens

import "github.com/jamestunnell/go-synth/slang"

type Func struct{}

const StrFUNC = "func"

func FUNC() slang.TokenInfo               { return &Func{} }
func (t *Func) Type() slang.TokenType { return slang.TokenFUNC }
func (t *Func) Value() string         { return StrFUNC }
