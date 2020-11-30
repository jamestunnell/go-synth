package complexslice

// Map maps the given values to a new slice using the given map function.
func Map(vals []complex128, mapFunc func(complex128) complex128) []complex128 {
	newVals := make([]complex128, len(vals))

	for i := 0; i < len(vals); i++ {
		newVals[i] = mapFunc(vals[i])
	}

	return newVals
}
