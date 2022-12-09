package synth

type Const[T any] struct {
	Val T
	Out *TypedOutput[T]
}

func NewConst[T any](val T) *Const[T] {
	return &Const[T]{
		Val: val,
		Out: NewTypedOutput[T](),
	}
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
