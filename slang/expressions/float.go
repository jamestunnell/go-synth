package expressions

type Float struct {
	val float64
}

func NewFloat(val float64) *Float {
	return &Float{val: val}
}
