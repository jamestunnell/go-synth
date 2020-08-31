package generators

import "github.com/jamestunnell/go-synth/node"

type ArrayType int

const (
	OneShot ArrayType = iota
	Repeat
)

type Array struct {
	Values []float64
	Type   ArrayType

	i int
}

func (a *Array) GetInterface() *node.Interface {
	return &node.Interface{
		Parameters: map[string]*node.ParamInfo{},
		Inputs:     []string{},
	}
}

func (a *Array) Initialize(srate float64) {
	a.i = 0

	if len(a.Values) == 0 {
		panic("array has no values")
	}
}

func (a *Array) Configure() {
}

func (a *Array) NumValuesLeft() int {
	return (len(a.Values) - a.i)
}

func (a *Array) Sample(out *node.Buffer) {
	switch a.Type {
	case OneShot:
		numLeft := a.NumValuesLeft()
		if numLeft >= out.Length {
			for i := 0; i < out.Length; i++ {
				out.Values[i] = a.Values[a.i]
				a.i++
			}
		} else {
			for i := 0; i < numLeft; i++ {
				out.Values[i] = a.Values[a.i]
				a.i++
			}

			for i := numLeft; i < out.Length; i++ {
				out.Values[i] = 0.0
			}
		}
	case Repeat:
		n := len(a.Values)
		for i := 0; i < out.Length; i++ {
			out.Values[i] = a.Values[a.i%n]
			a.i++
		}
	}
}
