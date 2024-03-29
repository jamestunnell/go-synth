package synth

import "reflect"

type Param interface {
	Type() string
	SetValue(any) error
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

func NewUint64Param(startVal uint64) *TypedParam[uint64] {
	return NewTypedParam(startVal)
}

func NewInt64Param(startVal int64) *TypedParam[int64] {
	return NewTypedParam(startVal)
}

func NewFloat64Param(startVal float64) *TypedParam[float64] {
	return NewTypedParam(startVal)
}

func NewBoolParam(startVal bool) *TypedParam[bool] {
	return NewTypedParam(startVal)
}

func NewStringParam(startVal string) *TypedParam[string] {
	return NewTypedParam(startVal)
}

func (tp *TypedParam[T]) Type() string {
	return reflect.TypeOf(tp.Value).String()
}

func (tp *TypedParam[T]) SetValue(val any) error {
	val2, ok := val.(T)
	if !ok {
		return NewErrTypeMismatch(
			reflect.TypeOf(tp.Value).String(),
			reflect.TypeOf(val).String(),
		)
	}

	tp.Value = val2

	return nil
}

func (tp *TypedParam[T]) GetValue() any {
	return tp.Value
}
