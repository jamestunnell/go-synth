package slang

type StatementType int

type Statement interface {
	Type() StatementType
	Equal(Statement) bool
}

const (
	StatementASSIGN StatementType = iota
	StatementIF
	StatementRETURN
)

func (st StatementType) String() string {
	var str string

	switch st {
	case StatementASSIGN:
		str = "ASSIGN"
	case StatementIF:
		str = "IF"
	case StatementRETURN:
		str = "RETURN"
	}

	return str
}
