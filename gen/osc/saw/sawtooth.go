package saw

import (
	"math"

	"github.com/jamestunnell/go-synth/gen/osc"
	"github.com/jamestunnell/go-synth/node"
)

const (
	oneOverPi = 1.0 / math.Pi
)

type Sawtooth struct {
	*osc.Osc
}

func init() {
	node.WorkingRegistry().RegisterCore(New())
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
