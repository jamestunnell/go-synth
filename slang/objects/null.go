package objects

import "github.com/jamestunnell/go-synth/slang"

type Null struct {
}

const StrNULL = "null"

func (obj *Null) Inspect() string {
	return StrNULL
}

func (obj *Null) Type() slang.ObjectType {
	return slang.ObjectNULL
}
