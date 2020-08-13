package constraints

var typeStrEqual = "Equal"

type Equal struct {
	SingleValue
}

func NewEqual(val float64) *Equal {
	return &Equal{SingleValue{Value: val, Type: typeStrEqual}}
}

func (eq *Equal) Allows(val float64) bool {
	return val == eq.Value
}
