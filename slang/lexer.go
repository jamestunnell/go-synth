package slang

type Lexer interface {
	NextToken() *Token
}
