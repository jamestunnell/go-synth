package unit

import (
	"github.com/jamestunnell/go-synth/node"
	"github.com/jamestunnell/go-synth/unit/gen/osc/saw"
	"github.com/jamestunnell/go-synth/unit/gen/osc/sine"
	"github.com/jamestunnell/go-synth/unit/gen/osc/square"
	"github.com/jamestunnell/go-synth/unit/gen/osc/triangle"
	"github.com/jamestunnell/go-synth/unit/proc/math/abs"
	"github.com/jamestunnell/go-synth/unit/proc/math/add"
	"github.com/jamestunnell/go-synth/unit/proc/math/div"
	"github.com/jamestunnell/go-synth/unit/proc/math/mul"
	"github.com/jamestunnell/go-synth/unit/proc/math/neg"
	"github.com/jamestunnell/go-synth/unit/proc/math/pow"
	"github.com/jamestunnell/go-synth/unit/proc/math/sub"
)

// RegisterBuiltinGenerators registers the core types for built-in unit
// generators in the given registry.
func RegisterBuiltinGenerators(reg *node.CoreRegistry) {
	// unit generators
	// reg.Register(oneshot.New())
	// reg.Register(repeat.New())
	reg.Register(saw.New())
	reg.Register(sine.New())
	reg.Register(square.New())
	reg.Register(triangle.New())
}

// RegisterBuiltinProcessors registers the core types for built-in unit
// processors in the given registry.
func RegisterBuiltinProcessors(reg *node.CoreRegistry) {
	// unit processors
	reg.Register(abs.New())
	reg.Register(add.New())
	reg.Register(div.New())
	reg.Register(mul.New())
	reg.Register(neg.New())
	reg.Register(pow.New())
	reg.Register(sub.New())
}
