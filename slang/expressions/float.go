package expressions

import "strconv"

type Float struct {
	val float64
}

func NewFloat(val float64) *Float {
	return &Float{val: val}
}

func (f *Float) String() string {
	return strconv.FormatFloat(f.val, 'g', -1, 64)
}
