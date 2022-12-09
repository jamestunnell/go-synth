package synth

type Const[T any] struct {
	Val T
	Out *TypedOutput[T]
}

func NewConstDeferVal[T any]() *Const[T] {
	return &Const[T]{
		Out: NewTypedOutput[T](),
	}
}

func NewConst[T any](val T) *Const[T] {
	return &Const[T]{
		Val: val,
		Out: NewTypedOutput[T](),
	}
}

func NewFloat64Const() Block {
	return NewConstDeferVal[float64]()
}

func NewInt64Const() Block {
	return NewConstDeferVal[int64]()
}

func NewUint64Const() Block {
	return NewConstDeferVal[uint64]()
}

func NewStringConst() Block {
	return NewConstDeferVal[string]()
}

func NewBoolConst() Block {
	return NewConstDeferVal[bool]()
}

func (cb *Const[T]) Initialize(srate float64, outDepth int) error {
	cb.Out.Initialize(outDepth)

	for i := 0; i < len(cb.Out.Buffer); i++ {
		cb.Out.Buffer[i] = cb.Val
	}

	return nil
}

func (cb *Const[T]) Configure() {

}

func (cb *Const[T]) Run() {

}
