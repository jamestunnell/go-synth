package synth

import "reflect"

type Output interface {
	Type() string
	Initialize(len int)
	Data() any
}

type TypedOutput[T any] struct {
	data []T
}

func NewTypedOutput[T any]() *TypedOutput[T] {
	return &TypedOutput[T]{
		data: []T{},
	}
}

func NewUint64Output() Output {
	return NewTypedOutput[uint64]()
}

func NewInt64Output() Output {
	return NewTypedOutput[int64]()
}

func NewFloat64Output() Output {
	return NewTypedOutput[float64]()
}

func NewBoolOutput() Output {
	return NewTypedOutput[bool]()
}

func NewStringOutput() Output {
	return NewTypedOutput[string]()
}

func (to *TypedOutput[T]) Type() string {
	var val T

	return reflect.TypeOf(val).String()
}

func (to *TypedOutput[T]) Initialize(len int) {
	to.data = make([]T, len)
}

func (to *TypedOutput[T]) Data() any {
	return to.data
}
