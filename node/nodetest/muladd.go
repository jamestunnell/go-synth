package nodetest

import "github.com/jamestunnell/go-synth/node"

type MulAdd struct {
	inBuf, mulKBuf, addKBuf *node.Buffer
	addK, mulK              float64
}

const (
	MulAddDefaultAddK = 0.0
	MulAddDefaultMulK = 1.0
)

func (mulAdd *MulAdd) Interface() *node.Interface {
	return &node.Interface{
		InputNames: []string{"In"},
		ControlDefaults: map[string]float64{
			"AddK": MulAddDefaultAddK,
			"MulK": MulAddDefaultMulK,
		},
	}
}

func (mulAdd *MulAdd) Initialize(srate float64, inputs, controls node.Map) {
	mulAdd.inBuf = inputs["In"].Output()
	mulAdd.mulKBuf = controls["MulK"].Output()
	mulAdd.addKBuf = controls["AddK"].Output()
}

func (mulAdd *MulAdd) Configure() {
	mulAdd.mulK = mulAdd.mulKBuf.Values[0]
	mulAdd.addK = mulAdd.addKBuf.Values[0]
}

func (mulAdd *MulAdd) Run(out *node.Buffer) {
	for i := 0; i < out.Length; i++ {
		out.Values[i] = (mulAdd.inBuf.Values[i] * mulAdd.mulK) + mulAdd.addK
	}
}
