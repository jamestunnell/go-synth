package complexslice

// FromFloats makes a complex slice from the given floats.
func FromFloats(input []float64) []complex128 {
	n := len(input)
	output := make([]complex128, n)

	for i := 0; i < n; i++ {
		output[i] = complex(input[i], 0.0)
	}

	return output
}
