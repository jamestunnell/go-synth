package unit

type Buffer struct {
	Values []float64
	Length int
}

func NewBuffer(n int) *Buffer {
	return &Buffer{Values: make([]float64, n), Length: n}
}

func (b *Buffer) Fill(val float64) {
	for i := 0; i < b.Length; i++ {
		b.Values[i] = val
	}
}
