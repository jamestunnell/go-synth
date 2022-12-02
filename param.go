package synth

import "reflect"

type Param interface {
	Type() string
	SetValue(any) bool
	GetValue() any
}

type TypedParam[T any] struct {
	Value T
}

func NewTypedParam[T any](startVal T) *TypedParam[T] {
	return &TypedParam[T]{
		Value: startVal,
	}
}

func NewUint64Param(startVal uint64) Param {
	return NewTypedParam(startVal)
}

func NewInt64Param(startVal int64) Param {
	return NewTypedParam(startVal)
}

func NewFloat64Param(startVal float64) Param {
	return NewTypedParam(startVal)
}

func NewBoolParam(startVal bool) Param {
	return NewTypedParam(startVal)
}

func NewStringParam(startVal string) Param {
	return NewTypedParam(startVal)
}

func (tp *TypedParam[T]) Type() string {
	return reflect.TypeOf(tp.Value).String()
}

func (tp *TypedParam[T]) SetValue(val any) bool {
	val2, ok := val.(T)
	if !ok {
		return false
	}

	tp.Value = val2

	return true
}

func (tp *TypedParam[T]) GetValue() any {
	return tp.Value
}
