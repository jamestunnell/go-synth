package unit

type Constraint interface {
	Allows(val float64) bool
}
