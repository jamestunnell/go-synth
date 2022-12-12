package tokens

import "github.com/jamestunnell/go-synth/slang"

type Comma struct{}

func COMMA() slang.Token               { return &Comma{} }
func (t *Comma) Type() slang.TokenType { return slang.TokenCOMMA }
func (t *Comma) Value() string         { return "," }
