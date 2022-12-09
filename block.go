package synth

type Block interface {
	Initialize(srate float64, outDepth int) error
	Configure()
	Run()
}

func BlockInterface(b Block) *Interface {
	ifc := NewInterface()

	ifc.Extract(b)

	return ifc
}
