package slang

type TokenType int

type Token interface {
	Type() string
	String() string
}
