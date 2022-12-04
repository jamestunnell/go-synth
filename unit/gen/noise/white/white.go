package white

import (
	"math/rand"
	"time"

	"github.com/jamestunnell/go-synth"
	"github.com/jamestunnell/go-synth/node"
	"github.com/jamestunnell/go-synth/node/mod"
	"github.com/jamestunnell/go-synth/util/param"
)

// White produces white noise (uncorrelated values) with a
// pseudo random number generator. Output is from -1 to 1.
type White struct {
	Seed *synth.TypedParam[int64]
	Out  *synth.TypedOutput[float64]

	outBuf []float64
	rnd    *rand.Rand
}

// ParamNameSeed is the name of the seed param
const ParamNameSeed = "Seed"

func ParamMods(seed int64) []node.Mod {
	return []node.Mod{mod.Param(ParamNameSeed, param.NewInt(seed))}
}

// New makes a new white noise block.
func New() *White {
	seed := time.Now().UnixNano()

	wh := &White{
		Seed: synth.NewInt64Param(seed),
	}

	wh.Out = synth.NewFloat64Output(wh)

	return wh
}

// Initialize initializes the block.
// Creates a new pseudo-RNG with the seed param.
func (w *White) Initialize(srate float64, outDepth int) error {
	w.Out.Initialize(outDepth)

	w.outBuf = w.Out.Buffer().([]float64)

	w.rnd = rand.New(rand.NewSource(w.Seed.Value))

	// log.Debug().Int64("seed", w.Seed.Value).Msg("new white noise rand")

	return nil
}

// Configure does nothing.
func (w *White) Configure() {
}

// Run generates white noise in the range [-1.0,1.0).
func (w *White) Run() {
	for i := 0; i < len(w.outBuf); i++ {
		w.outBuf[i] = (w.rnd.Float64() * 2.0) - 1.0
	}
}
