package triangle

import (
	"math"

	"github.com/jamestunnell/go-synth/gen/osc"
	"github.com/jamestunnell/go-synth/node"
)

const (
	twoOverPi = 2.0 / math.Pi
)

func New(params *osc.Params) node.Node {
	return osc.New(params, triangleWave)
}

func triangleWave(phase float64) float64 {
	return math.Abs(twoOverPi*phase) - 1.0
}
