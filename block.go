package synth

import (
	"path"
	"reflect"
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

// BlockPath returns the full block path, including package path and block struct Name.
func BlockPath(c Block) string {
	t := reflect.TypeOf(c).Elem()

	return path.Join(t.PkgPath(), t.Name())
}

func BlockMaker[T Block](f func() T) MakeBlockFunc {
	return func() Block {
		return f()
	}
}
