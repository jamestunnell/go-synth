package synth

type ConstBlock[T any] struct {
	Value  T
	Output *TypedOutput[T]
}

func NewConstBlock[T any](val T) *ConstBlock[T] {
	return &ConstBlock[T]{
		Value:  val,
		Output: NewTypedOutput[T](),
	}
}

func (cb *ConstBlock[T]) Initialize() error {
	return nil
}

func (cb *ConstBlock[T]) Configure() {

}

func (cb *ConstBlock[T]) Run() {

}
