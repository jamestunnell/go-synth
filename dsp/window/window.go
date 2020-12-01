package window

// Window makes windows of a certain kind.
type WindowMaker interface {
	// Make window of given size.
	Make(size int) []float64
}
