package array

import "github.com/jamestunnell/go-synth/node"

type repeat struct {
	*array
}

func Repeat(vals []float64) node.Node {
	o := &repeat{}
	o.array = newArray(vals)
	return o
}

func (r *repeat) Run() {
	for i := 0; i < r.array.outBuf.Length; i++ {
		r.array.outBuf.Values[i] = r.array.vals[r.array.idx%r.array.numVals]
		r.array.idx++
	}
}
