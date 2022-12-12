package expressions

import "github.com/jamestunnell/go-synth/slang"

type Identifier struct {
	Name string
}

func NewIdentifier(name string) *Identifier {
	return &Identifier{Name: name}
}

func (i *Identifier) Equal(other slang.Expression) bool {
	i2, ok := other.(*Identifier)
	if !ok {
		return false
	}

	return i2.Name == i.Name
}
