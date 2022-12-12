package tokens

import "github.com/jamestunnell/go-synth/slang"

type Slash struct{}

const (
	StrSLASH = "/"
)

func SLASH() slang.Token               { return &Slash{} }
func (t *Slash) Type() slang.TokenType { return slang.TokenSLASH }
func (t *Slash) Value() string         { return StrSLASH }
