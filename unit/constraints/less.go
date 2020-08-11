package constraints

const typeStrLess = "Less"

type Less struct {
	SingleValue
}

func NewLess(val float64) *Less {
	return &Less{SingleValue{Value: val, Type: typeStrLess}}
}

func (less *Less) Allows(val float64) bool {
	return val < less.SingleValue.Value
}
