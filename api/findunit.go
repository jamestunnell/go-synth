package api

import (
	"reflect"

	"github.com/jamestunnell/go-synth/node"
)

func findUnit(name string, cores []node.Core) node.Core {
	for _, core := range cores {
		if reflect.TypeOf(core).Elem().Name() == name {
			return core
		}
	}

	return nil
}
