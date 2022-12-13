package tokens

import "github.com/jamestunnell/go-synth/slang"

type False struct{}

const StrFALSE = "false"

func FALSE() slang.TokenInfo           { return &False{} }
func (t *False) Type() slang.TokenType { return slang.TokenFALSE }
func (t *False) Value() string         { return StrFALSE }
