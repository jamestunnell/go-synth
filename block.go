package synth

import (
	"reflect"

	"github.com/jamestunnell/go-synth/util/typeregistry"
)

type Block interface {
	Initialize(srate float64, outDepth int) error
	Configure()
	Run()
}

func BlockInterface(b Block) *Interface {
	ifc := NewInterface()

	ifc.Extract(b)

	return ifc
}

// BlockPath returns the full core path, including package path and core type.
func BlockPath(c Block) string {
	return typeregistry.TypePath(reflect.TypeOf(c).Elem())
}

func BlockMaker[T Block](f func() T) MakeBlockFunc {
	return func() Block {
		return f()
	}
}
