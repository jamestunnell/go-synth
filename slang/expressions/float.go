package expressions

import (
	"github.com/jamestunnell/go-synth/slang"
)

type Float struct {
	val float64
}

func NewFloat(val float64) *Float {
	return &Float{val: val}
}

// func (f *Float) String() string {
// 	return strconv.FormatFloat(f.val, 'g', -1, 64)
// }

func (f *Float) Equal(other slang.Expression) bool {
	f2, ok := other.(*Float)
	if !ok {
		return false
	}

	return f2.val == f.val
}
