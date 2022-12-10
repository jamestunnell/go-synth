package synth

import "reflect"

type Output interface {
	Type() string
	Initialize(len int)
}

type TypedOutput[T any] struct {
	Buffer []T
}

func NewTypedOutput[T any]() *TypedOutput[T] {
	return &TypedOutput[T]{
		Buffer: []T{},
	}
}

func NewUint64Output() *TypedOutput[uint64] {
	return NewTypedOutput[uint64]()
}

func NewInt64Output() *TypedOutput[int64] {
	return NewTypedOutput[int64]()
}

func NewFloat64Output() *TypedOutput[float64] {
	return NewTypedOutput[float64]()
}

func NewBoolOutput() *TypedOutput[bool] {
	return NewTypedOutput[bool]()
}

func NewStringOutput() *TypedOutput[string] {
	return NewTypedOutput[string]()
}

func (to *TypedOutput[T]) Type() string {
	var val T

	return reflect.TypeOf(val).String()
}

func (to *TypedOutput[T]) Initialize(len int) {
	to.Buffer = make([]T, len)
}
