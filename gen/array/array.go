package array

import "github.com/jamestunnell/go-synth/node"

type array struct {
	vals []float64

	outBuf  *node.Buffer
	idx     int
	numVals int
}

func newArray(vals []float64) *array {
	return &array{vals: vals}
}

func (a *array) Buffer() *node.Buffer {
	return a.outBuf
}

func (a *array) Controls() map[string]node.Node {
	return map[string]node.Node{}
}

func (a *array) Inputs() map[string]node.Node {
	return map[string]node.Node{}
}

func (a *array) Initialize(srate float64, depth int) {
	if len(a.vals) == 0 {
		panic("array has no values")
	}

	a.outBuf = node.NewBuffer(depth)
	a.numVals = len(a.vals)
	a.idx = 0
}

func (a *array) Configure() {
}
