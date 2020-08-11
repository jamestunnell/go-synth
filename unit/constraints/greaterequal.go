package constraints

const typeStrGreaterEqual = "GreaterEqual"

type GreaterEqual struct {
	SingleValue
}

func NewGreaterEqual(val float64) *GreaterEqual {
	return &GreaterEqual{SingleValue{Value: val, Type: typeStrGreaterEqual}}
}

func (ge *GreaterEqual) Allows(val float64) bool {
	return val >= ge.Value
}
