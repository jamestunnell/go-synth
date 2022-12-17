package slang

type Object interface {
	Type() ObjectType
	Inspect() string
}

type ObjectType int

const (
	ObjectBOOL ObjectType = iota
	ObjectFLOAT
	ObjectINTEGER
	ObjectNULL
)

func (ot ObjectType) String() string {
	var str string

	switch ot {
	case ObjectINTEGER:
		str = "INTEGER"
	case ObjectBOOL:
		str = "BOOL"
	case ObjectNULL:
		str = "NULL"
	}
	return str
}
