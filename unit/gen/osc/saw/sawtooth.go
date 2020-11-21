package saw

import (
	"math"

	"github.com/jamestunnell/go-synth/node"
	"github.com/jamestunnell/go-synth/unit/gen/osc"
)

const (
	oneOverPi = 1.0 / math.Pi
)

type Sawtooth struct {
	*osc.Osc
}

func NewNode(freq, phase *node.Node) *node.Node {
	return osc.NewNode(New(), freq, phase)
}

func New() *Sawtooth {
	s := &Sawtooth{}
	s.Osc = &osc.Osc{}
	return s
}

func (s *Sawtooth) Run(out *node.Buffer) {
	s.Osc.Run(sawtoothWave, out)
}

func sawtoothWave(phase float64) float64 {
	return phase * oneOverPi
}
