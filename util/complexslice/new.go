package complexslice

// New makes a new slice, initializing each value using the given func.
func New(n int, f func(idx int) complex128) []complex128 {
	vals := make([]complex128, n)

	for i := 0; i < n; i++ {
		vals[i] = f(i)
	}

	return vals
}
