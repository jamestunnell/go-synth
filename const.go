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

func NewConstFromAny(val any) (Block, error) {
	var blk Block

	switch vv := val.(type) {
	case float64:
		blk = NewConst(vv)
	case int64:
		blk = NewConst(vv)
	case uint64:
		blk = NewConst(vv)
	case string:
		blk = NewConst(vv)
	case bool:
		blk = NewConst(vv)
	}

	if blk == nil {
		return nil, NewErrUnsupportedType(val)
	}

	return blk, nil
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
