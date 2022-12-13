package tokens

import "github.com/jamestunnell/go-synth/slang"

type Eof struct{}

func EOF() slang.TokenInfo           { return &Eof{} }
func (t *Eof) Type() slang.TokenType { return slang.TokenEOF }
func (t *Eof) Value() string         { return "" }
