package expressions

type Identifier struct {
	Name string
}

func NewIdentifier(name string) *Identifier {
	return &Identifier{Name: name}
}

func (i *Identifier) String() string {
	return i.Name
}
