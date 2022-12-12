package expressions

import (
	"github.com/jamestunnell/go-synth/slang"
)

type Integer struct {
	val int64
}

func NewInteger(val int64) *Integer {
	return &Integer{val: val}
}

// func (i *Integer) String() string {
// 	return strconv.FormatInt(i.val, 10)
// }

func (i *Integer) Equal(other slang.Expression) bool {
	i2, ok := other.(*Integer)
	if !ok {
		return false
	}

	return i2.val == i.val
}
