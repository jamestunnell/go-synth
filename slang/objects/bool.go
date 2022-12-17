package objects

import (
	"strconv"

	"github.com/jamestunnell/go-synth/slang"
)

type Bool struct {
	Value bool
}

func NewBool(val bool) slang.Object {
	return &Bool{Value: val}
}

func (obj *Bool) Inspect() string {
	return strconv.FormatBool(obj.Value)
}

func (obj *Bool) Type() slang.ObjectType {
	return slang.ObjectBOOL
}
