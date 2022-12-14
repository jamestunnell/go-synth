package slang

type ExprType int

type Expression interface {
	Type() ExprType
	Equal(Expression) bool
}

const (
	ExprADD ExprType = iota
	ExprANONYMOUSFUNC
	ExprBOOL
	ExprDIVIDE
	ExprEQUAL
	ExprFLOAT
	ExprGREATER
	ExprGREATEREQUAL
	ExprIDENTIFIER
	ExprIF
	ExprINTEGER
	ExprLESS
	ExprLESSEQUAL
	ExprMULTIPLY
	ExprNEGATIVE
	ExprNOT
	ExprNOTEQUAL
	ExprSUBTRACT
)

func (st ExprType) String() string {
	var str string

	switch st {
	case ExprADD:
		str = "ADD"
	case ExprANONYMOUSFUNC:
		str = "ANONYMOUSFUNC"
	case ExprBOOL:
		str = "BOOL"
	case ExprDIVIDE:
		str = "DIVIDE"
	case ExprEQUAL:
		str = "EQUAL"
	case ExprFLOAT:
		str = "FLOAT"
	case ExprGREATER:
		str = "GREATER"
	case ExprGREATEREQUAL:
		str = "GREATEREQUAL"
	case ExprIDENTIFIER:
		str = "IDENTIFIER"
	case ExprIF:
		str = "IF"
	case ExprINTEGER:
		str = "INTEGER"
	case ExprLESS:
		str = "LESS"
	case ExprLESSEQUAL:
		str = "LESSEQUAL"
	case ExprMULTIPLY:
		str = "MULTIPLY"
	case ExprNEGATIVE:
		str = "NEGATIVE"
	case ExprNOT:
		str = "NOT"
	case ExprNOTEQUAL:
		str = "NOTEQUAL"
	case ExprSUBTRACT:
		str = "SUBTRACT"
	}

	return str
}
