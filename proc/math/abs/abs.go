package abs

import (
	"math"

	"github.com/jamestunnell/go-synth/node"
)

type Abs struct {
	inBuf *node.Buffer
}

func init() {
	node.WorkingRegistry().RegisterCore(New())
}

func NewNode(in *node.Node) *node.Node {
	inputs := node.Map{"In": in}
	return node.NewNode(New(), inputs, node.Map{})
}

func New() *Abs {
	return &Abs{}
}

func (a *Abs) Interface() *node.Interface {
	return &node.Interface{
		InputNames:      []string{"In"},
		ControlDefaults: map[string]float64{},
	}
}

func (a *Abs) Initialize(srate float64, inputs, controls node.Map) {
	a.inBuf = inputs["In"].Output()
}

func (a *Abs) Configure() {
}

func (a *Abs) Run(out *node.Buffer) {
	for i := 0; i < out.Length; i++ {
		out.Values[i] = math.Abs(a.inBuf.Values[i])
	}
}
