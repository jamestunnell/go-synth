package square

import (
	"github.com/jamestunnell/go-synth/gen/osc"
	"github.com/jamestunnell/go-synth/node"
)

type Square struct {
	*osc.Osc
}

func NewNode(freq, phase *node.Node) *node.Node {
	return osc.NewNode(New(), freq, phase)
}

func New() *Square {
	s := &Square{}
	s.Osc = &osc.Osc{}
	return s
}

func (s *Square) Run(out *node.Buffer) {
	s.Osc.Run(squareWave, out)
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
