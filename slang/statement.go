package slang

type Statement interface {
	Type() string
	Equal(Statement) bool
}
