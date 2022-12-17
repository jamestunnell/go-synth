package objects

import (
	"strconv"

	"github.com/jamestunnell/go-synth/slang"
)

type Integer struct {
	Value int64
}

func NewInteger(val int64) slang.Object {
	return &Integer{Value: val}
}

func (obj *Integer) Inspect() string {
	return strconv.FormatInt(obj.Value, 10)
}

func (obj *Integer) Type() slang.ObjectType {
	return slang.ObjectINTEGER
}
