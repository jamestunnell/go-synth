package proc

import (
	"github.com/jamestunnell/go-synth/node"
	"github.com/jamestunnell/go-synth/unit/proc/math"
)

// RegisterBuiltin registers the core types for built-in unit
// processors in the given registry.
func RegisterBuiltin(reg *node.CoreRegistry) {
	// unit processors
	reg.Register(&math.Abs{})
	reg.Register(&math.Add{})
	reg.Register(&math.Div{})
	reg.Register(&math.Mul{})
	reg.Register(&math.Neg{})
	reg.Register(&math.Pow{})
	reg.Register(&math.Sub{})
}
