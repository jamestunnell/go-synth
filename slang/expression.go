package slang

type Expression interface {
	Equal(Expression) bool
}
