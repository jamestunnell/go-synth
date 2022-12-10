package proc

import (
	"github.com/jamestunnell/go-synth"
	"github.com/jamestunnell/go-synth/unit/proc/math"
)

// RegisterBuiltin registers the core types for built-in unit
// processors in the given registry.
func RegisterBuiltin(reg *synth.BlockRegistry) {
	// unit processors
	reg.Register(synth.BlockMaker(math.NewAbs))
	reg.Register(synth.BlockMaker(math.NewAdd))
	reg.Register(synth.BlockMaker(math.NewDiv))
	reg.Register(synth.BlockMaker(math.NewMul))
	reg.Register(synth.BlockMaker(math.NewNeg))
	reg.Register(synth.BlockMaker(math.NewPow))
	reg.Register(synth.BlockMaker(math.NewSub))
}
