package synth

type Block interface {
	Initialize(srate float64, outDepth int) error
	Configure()
	Run()
}
