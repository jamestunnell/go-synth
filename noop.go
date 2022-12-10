package synth

// NoOp implements the Block interface, but does nothing
type NoOp struct {
}

func (no *NoOp) Initialize(srate float64, outDepth int) error {
	return nil
}

func (no *NoOp) Configure() {
}

func (no *NoOp) Run() {
}
