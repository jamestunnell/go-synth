package synth

type ConstBlock[T any] struct {
	Val T
	Out *TypedOutput[T]
}

func NewConstBlock[T any](val T) *ConstBlock[T] {
	cb := &ConstBlock[T]{
		Val: val,
	}

	cb.Out = NewTypedOutput[T](cb)

	return cb
}

func (cb *ConstBlock[T]) Initialize(srate float64, outDepth int) error {
	cb.Out.Initialize(outDepth)

	for i := 0; i < len(cb.Out.BufferValues); i++ {
		cb.Out.BufferValues[i] = cb.Val
	}

	return nil
}

func (cb *ConstBlock[T]) Configure() {

}

func (cb *ConstBlock[T]) Run() {

}
