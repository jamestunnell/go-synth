package triangle

import (
	"math"

	"github.com/jamestunnell/go-synth/gen/osc"
	"github.com/jamestunnell/go-synth/node"
)

const (
	twoOverPi = 2.0 / math.Pi
)

type Triangle struct {
	*osc.Osc
}

func init() {
	node.WorkingRegistry().RegisterCore(New())
}

func NewNode(freq, phase *node.Node) *node.Node {
	return osc.NewNode(New(), freq, phase)
}

func New() *Triangle {
	s := &Triangle{}
	s.Osc = &osc.Osc{}
	return s
}

func (s *Triangle) Run(out *node.Buffer) {
	s.Osc.Run(triangleWave, out)
}

func triangleWave(phase float64) float64 {
	return math.Abs(twoOverPi*phase) - 1.0
}
