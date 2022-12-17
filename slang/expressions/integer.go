package expressions

import (
	"github.com/jamestunnell/go-synth/slang"
)

type Integer struct {
	Value int64
}

func NewInteger(val int64) *Integer {
	return &Integer{Value: val}
}

// func (i *Integer) String() string {
// 	return strconv.FormatInt(i.Value, 10)
// }

func (i *Integer) Type() slang.ExprType { return slang.ExprINTEGER }

func (i *Integer) Equal(other slang.Expression) bool {
	i2, ok := other.(*Integer)
	if !ok {
		return false
	}

	return i2.Value == i.Value
}
