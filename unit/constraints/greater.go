package constraints

const typeStrGreater = "Greater"

type Greater struct {
	SingleValue
}

func NewGreater(val float64) *Greater {
	return &Greater{SingleValue{Value: val, Type: typeStrGreater}}
}

func (greater *Greater) Allows(val float64) bool {
	return val > greater.Value
}
