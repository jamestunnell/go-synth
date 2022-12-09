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
func RegisterBuiltin(reg *synth.BlockRegistry) {
	// unit generators
	reg.Register(synth.BlockMaker(array.NewOneshotNoVals))
	reg.Register(synth.BlockMaker(array.NewRepeatNoVals))

	reg.Register(synth.BlockMaker(env.NewADSR))
	reg.Register(synth.BlockMaker(env.NewDecay))

	reg.Register(synth.BlockMaker(noise.NewBrown))
	reg.Register(synth.BlockMaker(noise.NewPink))
	reg.Register(synth.BlockMaker(noise.NewWhite))

	reg.Register(synth.BlockMaker(osc.NewSawtooth))
	reg.Register(synth.BlockMaker(osc.NewSine))
	reg.Register(synth.BlockMaker(osc.NewSquare))
	reg.Register(synth.BlockMaker(osc.NewTriangle))
}
