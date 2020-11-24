package array

// Fill fills all elements of the given slice with the given value.
func Fill(dst []float64, val float64) {
	for i := 0; i < len(dst); i++ {
		dst[i] = val
	}
}
