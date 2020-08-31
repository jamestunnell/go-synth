package node

type NoOpCore struct {
}

func (noop *NoOpCore) GetInterface() *Interface {
	return &Interface{
		Parameters: map[string]*ParamInfo{},
		Inputs:     []string{},
	}
}

func (noop *NoOpCore) Initialize(srate float64) {

}

func (noop *NoOpCore) Configure() {
}

func (noop *NoOpCore) Sample(out *Buffer) {
}
