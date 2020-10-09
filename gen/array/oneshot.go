package array

import "github.com/jamestunnell/go-synth/node"

type oneshot struct {
	*array
}

func OneShot(vals []float64) node.Node {
	o := &oneshot{}
	o.array = newArray(vals)
	return o
}

func (o *oneshot) Run() {
	for i := 0; i < o.array.outBuf.Length; i++ {
		var outVal float64
		if o.array.idx < o.array.numVals {
			outVal = o.array.vals[o.array.idx]
			o.array.idx++
		}
		o.array.outBuf.Values[i] = outVal
	}
}
