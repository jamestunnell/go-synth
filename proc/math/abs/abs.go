package abs

import (
	"math"

	"github.com/jamestunnell/go-synth/node"
)

type Abs struct {
	inBuf *node.Buffer
}

func NewNode(in *node.Node) *node.Node {
	inputs := map[string]*node.Node{"In": in}
	return node.New(New(), inputs, map[string]*node.Node{})
}

func New() *Abs {
	return &Abs{}
}

func (a *Abs) Initialize(srate float64, inputs, controls map[string]*node.Node) {
	a.inBuf = node.GetOutput(inputs, "In")
}

func (a *Abs) Configure() {
}

func (a *Abs) Run(out *node.Buffer) {
	for i := 0; i < out.Length; i++ {
		out.Values[i] = math.Abs(a.inBuf.Values[i])
	}
}
