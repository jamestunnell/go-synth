package unit

type Buffer struct {
	Values []float64
	Length int
}

func NewBuffer(n int) *Buffer {
	return &Buffer{Values: make([]float64, n), Length: n}
}
