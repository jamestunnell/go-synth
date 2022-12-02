package expressions

type Identifier struct {
	Name string
}

func NewIdentifier(name string) *Identifier {
	return &Identifier{Name: name}
}
