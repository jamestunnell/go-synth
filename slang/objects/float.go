package objects

import (
	"strconv"

	"github.com/jamestunnell/go-synth/slang"
)

type Float struct {
	Value float64
}

func NewFloat(val float64) slang.Object {
	return &Float{Value: val}
}

func (obj *Float) Inspect() string {
	return strconv.FormatFloat(obj.Value, 'g', -1, 64)
}

func (obj *Float) Type() slang.ObjectType {
	return slang.ObjectFLOAT
}
