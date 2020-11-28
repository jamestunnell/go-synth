package white

import (
	"math/rand"
	"time"

	"github.com/jamestunnell/go-synth/node"
	"github.com/jamestunnell/go-synth/node/mod"
	"github.com/jamestunnell/go-synth/util/param"
)

// White produces white noise (uncorrelated values) with a
// pseudo random number generator. Output is from -1 to 1.
type White struct {
	rnd *rand.Rand
}

// ParamNameSeed is the name of the seed param
const ParamNameSeed = "Seed"

// New makes a new White node
func New(moreMods ...node.Mod) *node.Node {
	seed := time.Now().UTC().UnixNano()

	return NewFromSeed(seed, moreMods...)
}

// NewFromSeed makes a new White node that uses the given seed.
func NewFromSeed(seed int64, moreMods ...node.Mod) *node.Node {
	mods := []node.Mod{mod.Param(ParamNameSeed, param.NewInt(seed))}

	mods = append(mods, moreMods...)

	return node.New(&White{}, mods...)
}

// Interface provides the node interface
func (w *White) Interface() *node.Interface {
	ifc := node.NewInterface()

	ifc.ParamTypes = map[string]param.Type{
		ParamNameSeed: param.Int,
	}

	return ifc
}

// Initialize initializes the node.
// Creates a new pseudo-RNG with the seed param.
func (w *White) Initialize(args *node.InitArgs) error {
	seed := args.Params[ParamNameSeed].Value().(int64)

	w.rnd = rand.New(rand.NewSource(seed))

	return nil
}

// Configure does nothing.
func (w *White) Configure() {
}

// Run runs the white noise generation process.
func (w *White) Run(out *node.Buffer) {
	for i := 0; i < out.Length; i++ {
		out.Values[i] = w.rnd.Float64()
	}
}
