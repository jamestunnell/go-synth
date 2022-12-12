package expressions

import "strconv"

type Integer struct {
	val int64
}

func NewInteger(val int64) *Integer {
	return &Integer{val: val}
}

func (i *Integer) String() string {
	return strconv.FormatInt(i.val, 10)
}
