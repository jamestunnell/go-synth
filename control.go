package synth

type Control interface {
	Input

	DefaultVal() any
}

type TypedControl[T any] struct {
	*TypedInput[T]

	defaultVal T
}

func NewTypedControl[T any](defaultVal T) *TypedControl[T] {
	return &TypedControl[T]{
		TypedInput: NewTypedInput[T](),
		defaultVal: defaultVal,
	}
}

func NewUint64Control(defaultVal uint64) *TypedControl[uint64] {
	return NewTypedControl(defaultVal)
}

func NewInt64Control(defaultVal int64) *TypedControl[int64] {
	return NewTypedControl(defaultVal)
}

func NewFloat64Control(defaultVal float64) *TypedControl[float64] {
	return NewTypedControl(defaultVal)
}

func NewBoolControl(defaultVal bool) *TypedControl[bool] {
	return NewTypedControl(defaultVal)
}

func NewStringControl(defaultVal string) *TypedControl[string] {
	return NewTypedControl(defaultVal)
}

func (ti *TypedControl[T]) DefaultVal() any {
	return ti.defaultVal
}
