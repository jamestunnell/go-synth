package synth

import "reflect"

type Input interface {
	Type() string
	Connected() bool
	Connect(Output) error
}

type TypedInput[T any] struct {
	Output Output
}

func NewTypedInput[T any]() *TypedInput[T] {
	return &TypedInput[T]{
		Output: nil,
	}
}

func NewUint64Input() *TypedInput[uint64] {
	return NewTypedInput[uint64]()
}

func NewInt64Input() *TypedInput[int64] {
	return NewTypedInput[int64]()
}

func NewFloat64Input() *TypedInput[float64] {
	return NewTypedInput[float64]()
}

func NewBoolInput() *TypedInput[bool] {
	return NewTypedInput[bool]()
}

func NewStringInput() *TypedInput[string] {
	return NewTypedInput[string]()
}

func (ti *TypedInput[T]) Type() string {
	var val T

	return reflect.TypeOf(val).String()
}

func (ti *TypedInput[T]) Connected() bool {
	return ti.Output != nil
}

func (ti *TypedInput[T]) Connect(out Output) error {
	expected := ti.Type()
	actual := out.Type()

	if expected != actual {
		return NewErrTypeMismatch(expected, actual)
	}

	ti.Output = out

	return nil
}
