package tokens

import "github.com/jamestunnell/go-synth/slang"

type Plus struct{}

const StrPLUS = "+"

func PLUS() slang.TokenInfo           { return &Plus{} }
func (t *Plus) Type() slang.TokenType { return slang.TokenPLUS }
func (t *Plus) Value() string         { return StrPLUS }
