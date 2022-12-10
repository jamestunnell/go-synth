package synth

type MonoTerminal struct {
	*NoOp

	In *TypedInput[float64]
}

type StereoTerminal struct {
	*NoOp

	In1 *TypedInput[float64]
	In2 *TypedInput[float64]
}

func IsTerminal(blk Block) bool {
	switch blk.(type) {
	case *MonoTerminal:
		return true
	case *StereoTerminal:
		return true
	}

	return false
}

func NewMonoTerminal() *MonoTerminal {
	return &MonoTerminal{
		NoOp: &NoOp{},
		In:   NewFloat64Input(),
	}
}

func NewStereoTerminal() *StereoTerminal {
	return &StereoTerminal{
		NoOp: &NoOp{},
		In1:  NewFloat64Input(),
		In2:  NewFloat64Input(),
	}
}
