package synth

type Control interface {
	Input

	MakeDefault() Block
}

type TypedControl[T any] struct {
	*TypedInput[T]

	DefaultVal T
}

func NewTypedControl[T any](defaultVal T) *TypedControl[T] {
	return &TypedControl[T]{
		TypedInput: NewTypedInput[T](),
		DefaultVal: defaultVal,
	}
}

func NewUint64Control(defaultVal uint64) Control {
	return NewTypedControl(defaultVal)
}

func NewInt64Control(defaultVal int64) Control {
	return NewTypedControl(defaultVal)
}

func NewFloat64Control(defaultVal float64) Control {
	return NewTypedControl(defaultVal)
}

func NewBoolControl(defaultVal bool) Control {
	return NewTypedControl(defaultVal)
}

func NewStringControl(defaultVal string) Control {
	return NewTypedControl(defaultVal)
}

func (ti *TypedControl[T]) MakeDefault() Block {
	return NewConstBlock(ti.DefaultVal)
}
