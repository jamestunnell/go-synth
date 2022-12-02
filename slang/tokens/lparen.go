package tokens

import "github.com/jamestunnell/go-synth/slang"

type LParen struct{}

const (
	StrLPAREN  = "("
	TypeLPAREN = "LPAREN"
)

func LPAREN() slang.Token       { return &LParen{} }
func (t *LParen) Type() string  { return TypeLPAREN }
func (t *LParen) Value() string { return StrLPAREN }
