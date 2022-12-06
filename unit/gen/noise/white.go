package noise

import (
	"math/rand"
	"time"

	"github.com/jamestunnell/go-synth"
)

// White produces white noise (uncorrelated values) with a
// pseudo random number generator. Output is from -1 to 1.
type White struct {
	Seed *synth.TypedParam[int64]
	Out  *synth.TypedOutput[float64]

	rnd *rand.Rand
}

// New makes a new white noise block.
func NewWhite() *White {
	seed := time.Now().UnixNano()

	return &White{
		Seed: synth.NewInt64Param(seed),
		Out:  synth.NewFloat64Output(),
	}
}

// Initialize initializes the block.
// Creates a new pseudo-RNG with the seed param.
func (w *White) Initialize(srate float64, outDepth int) error {
	w.Out.Initialize(outDepth)

	w.rnd = rand.New(rand.NewSource(w.Seed.Value))

	// log.Debug().Int64("seed", w.Seed.Value).Msg("new white noise rand")

	return nil
}

// Configure does nothing.
func (w *White) Configure() {
}

// Run generates white noise in the range [-1.0,1.0).
func (w *White) Run() {
	for i := 0; i < len(w.Out.BufferValues); i++ {
		w.Out.BufferValues[i] = (w.rnd.Float64() * 2.0) - 1.0
	}
}
