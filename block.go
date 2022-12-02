package synth

type Block interface {
	Initialize() error
	Configure()
	Run()
}
