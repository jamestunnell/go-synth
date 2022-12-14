package parser

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
