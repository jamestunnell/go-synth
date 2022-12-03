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

func NewUint64Output(parent Block) Output {
	return NewTypedOutput[uint64](parent)
}

func NewInt64Output(parent Block) Output {
	return NewTypedOutput[int64](parent)
}

func NewFloat64Output(parent Block) Output {
	return NewTypedOutput[float64](parent)
}

func NewBoolOutput(parent Block) Output {
	return NewTypedOutput[bool](parent)
}

func NewStringOutput(parent Block) Output {
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
