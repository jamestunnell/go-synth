package node

// Buffer is used to capture node output.
type Buffer struct {
	// Values are the buffer values
	Values []float64
	// Length is the buffer length
	Length int
}

// NewBuffer creates a buffer of the given size.
func NewBuffer(n int) *Buffer {
	return &Buffer{Values: make([]float64, n), Length: n}
}
