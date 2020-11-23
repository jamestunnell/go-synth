package gen

import (
	"github.com/jamestunnell/go-synth/node"
	"github.com/jamestunnell/go-synth/unit/gen/array"
	"github.com/jamestunnell/go-synth/unit/gen/osc"
)

// RegisterBuiltin registers the core types for built-in unit
// generators in the given registry.
func RegisterBuiltin(reg *node.CoreRegistry) {
	// unit generators
	reg.Register(&array.Array{})
	reg.Register(&osc.Osc{})
}
