package slang

type TokenType int

type Token interface {
	Type() string
	Value() string
}
