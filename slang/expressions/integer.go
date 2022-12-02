package expressions

type Integer struct {
	val int64
}

func NewInteger(val int64) *Integer {
	return &Integer{val: val}
}
