package gen

import (
	"github.com/jamestunnell/go-synth"
	"github.com/jamestunnell/go-synth/unit/gen/array"
	"github.com/jamestunnell/go-synth/unit/gen/env"
	"github.com/jamestunnell/go-synth/unit/gen/noise"
	"github.com/jamestunnell/go-synth/unit/gen/osc"
)

// RegisterBuiltin registers the core types for built-in unit
// generators in the given registry.
func RegisterBuiltin(reg synth.BlockRegistry) {
	// unit generators
	reg.Register(&array.Oneshot{})
	reg.Register(&array.Repeat{})

	reg.Register(&env.ADSR{})
	reg.Register(&env.Decay{})

	reg.Register(&noise.Brown{})
	reg.Register(&noise.Pink{})
	reg.Register(&noise.White{})

	reg.Register(&osc.Sawtooth{})
	reg.Register(&osc.Sine{})
	reg.Register(&osc.Square{})
	reg.Register(&osc.Triangle{})
}
