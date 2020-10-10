package sine

import (
	"math"

	"github.com/jamestunnell/go-synth/gen/osc"
	"github.com/jamestunnell/go-synth/node"
)

const (
	fourOverPi           = 4.0 / math.Pi
	negFourOverPiSquared = -4.0 / (math.Pi * math.Pi)
	kSineP               = 0.225
)

func New(params *osc.Params) node.Node {
	return osc.New(params, sineWave)
}

func sineWave(phase float64) float64 {
	y := fourOverPi*phase + negFourOverPiSquared*phase*math.Abs(phase)
	// for extra precision
	return kSineP*(y*math.Abs(y)-y) + y // Q * y + P * y * y.abs
}
