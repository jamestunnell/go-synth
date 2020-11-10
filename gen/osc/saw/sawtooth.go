package saw

import (
	"math"

	"github.com/jamestunnell/go-synth/gen/osc"
	"github.com/jamestunnell/go-synth/node"
)

const (
	oneOverPi = 1.0 / math.Pi
)

func NewNode(freq, phase *node.Node) *node.Node {
	return osc.NewNode(freq, phase, sawtoothWave)
}

func New(params *osc.Params) node.Node {
	return osc.New(params, sawtoothWave)
}

func sawtoothWave(phase float64) float64 {
	return phase * oneOverPi
}
