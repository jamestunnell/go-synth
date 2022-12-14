package parser

import (
	"github.com/jamestunnell/go-synth/slang"
)

type Precedence int

const (
	PrecedenceLOWEST      Precedence = iota
	PrecedenceEQUALS                 // ==
	PrecedenceLESSGREATER            // > or <
	PrecedenceSUM                    // +
	PrecedencePRODUCT                // *
	PrecedencePREFIX                 // -X or !X
	PrecedenceCALL                   // myFunction(X)
)

var precedences = map[slang.TokenType]Precedence{
	slang.TokenEQUAL:    PrecedenceEQUALS,
	slang.TokenNOTEQUAL: PrecedenceEQUALS,
	slang.TokenLESS:     PrecedenceLESSGREATER,
	slang.TokenGREATER:  PrecedenceLESSGREATER,
	slang.TokenPLUS:     PrecedenceSUM,
	slang.TokenMINUS:    PrecedenceSUM,
	slang.TokenSLASH:    PrecedencePRODUCT,
	slang.TokenSTAR:     PrecedencePRODUCT,
}
