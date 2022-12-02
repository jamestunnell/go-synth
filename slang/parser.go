package slang

type Parser interface {
	ParseProgram() (*Program, error)
}
