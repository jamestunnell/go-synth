package network

import (
	"reflect"

	"github.com/jamestunnell/go-synth"
)

type BlockMap map[string]synth.Block

func (m BlockMap) Equal(other BlockMap) bool {
	if len(m) != len(other) {
		return false
	}

	for name, b := range m {
		b2, found := other[name]
		if !found {
			return false
		}

		if synth.BlockPath(b) != synth.BlockPath(b2) {
			return false
		}

		pVals := synth.BlockInterface(b).ParamVals()
		pVals2 := synth.BlockInterface(b).ParamVals()
		for name, val := range pVals {
			val2, found := pVals2[name]
			if !found {
				return false
			}

			if !reflect.DeepEqual(val, val2) {
				return false
			}
		}
	}

	return true
}
