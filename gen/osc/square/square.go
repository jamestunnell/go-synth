package square

import (
	"github.com/jamestunnell/go-synth/gen/osc"
	"github.com/jamestunnell/go-synth/node"
)

func New(params *osc.Params) node.Node {
	return osc.New(params, squareWave)
}

func squareWave(phase float64) float64 {
	var y float64
	if phase >= 0.0 {
		y = 1.0
	} else {
		y = -1.0
	}

	return y
}
