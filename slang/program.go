package slang

type Program struct {
	Statements []Statement
}

func NewProgram(statements ...Statement) *Program {
	return &Program{
		Statements: statements,
	}
}

func (p *Program) AddStatement(s Statement) {
	p.Statements = append(p.Statements, s)
}
