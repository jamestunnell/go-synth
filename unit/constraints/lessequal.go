package constraints

const typeStrLessEqual = "LessEqual"

type LessEqual struct {
	SingleValue
}

func NewLessEqual(val float64) *LessEqual {
	return &LessEqual{SingleValue{Value: val, Type: typeStrLessEqual}}
}

func (le *LessEqual) Allows(val float64) bool {
	return val <= le.Value
}
