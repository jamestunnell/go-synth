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

	outBuf := cb.Out.Buffer().([]T)

	for i := 0; i < len(outBuf); i++ {
		outBuf[i] = cb.Val
	}

	return nil
}

func (cb *ConstBlock[T]) Configure() {

}

func (cb *ConstBlock[T]) Run() {

}
