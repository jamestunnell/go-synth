package array

func Fill(dst []float64, val float64) {
	for i := 0; i < len(dst); i++ {
		dst[i] = val
	}
}
