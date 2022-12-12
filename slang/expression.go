package slang

type ExprType int

type Expression interface {
	Type() ExprType
	Equal(Expression) bool
}

const (
	ExprADD ExprType = iota
	ExprANONYMOUSFUNC
	ExprDIVIDE
	ExprEQUAL
	ExprFLOAT
	ExprGREATER
	ExprGREATEREQUAL
	ExprIDENTIFIER
	ExprINTEGER
	ExprLESS
	ExprLESSEQUAL
	ExprMULTIPLY
	ExprNOTEQUAL
	ExprSUBTRACT
)
