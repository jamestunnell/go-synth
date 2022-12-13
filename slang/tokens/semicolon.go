package tokens

import "github.com/jamestunnell/go-synth/slang"

type Semicolon struct{}

func SEMICOLON() slang.TokenInfo           { return &Semicolon{} }
func (t *Semicolon) Type() slang.TokenType { return slang.TokenSEMICOLON }
func (t *Semicolon) Value() string         { return ";" }
