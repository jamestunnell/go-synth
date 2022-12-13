package tokens

import "github.com/jamestunnell/go-synth/slang"

type Return struct{}

const StrRETURN = "return"

func RETURN() slang.TokenInfo           { return &Return{} }
func (t *Return) Type() slang.TokenType { return slang.TokenRETURN }
func (t *Return) Value() string         { return StrRETURN }
