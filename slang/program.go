package slang

type Program struct {
	Statements []Statement
}

func NewProgram() *Program {
	return &Program{
		Statements: []Statement{},
	}
}

func (p *Program) AddStatement(s Statement) {
	p.Statements = append(p.Statements, s)
}
