package synth

import "reflect"

type Output interface {
	Parent() Block
	Type() string
	Initialize(len int)
	Buffer() any
}

type TypedOutput[T any] struct {
	parent Block
	buffer []T
}

func NewTypedOutput[T any](parent Block) *TypedOutput[T] {
	return &TypedOutput[T]{
		buffer: []T{},
		parent: parent,
	}
}

func NewUint64Output(parent Block) *TypedOutput[uint64] {
	return NewTypedOutput[uint64](parent)
}

func NewInt64Output(parent Block) *TypedOutput[int64] {
	return NewTypedOutput[int64](parent)
}

func NewFloat64Output(parent Block) *TypedOutput[float64] {
	return NewTypedOutput[float64](parent)
}

func NewBoolOutput(parent Block) *TypedOutput[bool] {
	return NewTypedOutput[bool](parent)
}

func NewStringOutput(parent Block) *TypedOutput[string] {
	return NewTypedOutput[string](parent)
}

func (to *TypedOutput[T]) Parent() Block {
	return to.parent
}

func (to *TypedOutput[T]) Type() string {
	var val T

	return reflect.TypeOf(val).String()
}

func (to *TypedOutput[T]) Initialize(len int) {
	to.buffer = make([]T, len)
}

func (to *TypedOutput[T]) Buffer() any {
	return to.buffer
}
