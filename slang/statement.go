package slang

type StatementType int

type Statement interface {
	Type() StatementType
	Equal(Statement) bool
}

const (
	StatementASSIGN StatementType = iota
	StatementBLOCK
	StatementRETURN
	StatementFUNC
	StatementEXPRESSION
)

func (st StatementType) String() string {
	var str string

	switch st {
	case StatementASSIGN:
		str = "ASSIGN"
	case StatementBLOCK:
		str = "BLOCK"
	case StatementEXPRESSION:
		str = "EXPRESSION"
	case StatementFUNC:
		str = "FUNC"
	case StatementRETURN:
		str = "RETURN"
	}

	return str
}
