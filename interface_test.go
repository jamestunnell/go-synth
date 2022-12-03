package synth_test

import (
	"testing"

	"github.com/jamestunnell/go-synth"
	"github.com/stretchr/testify/assert"
)

type TestBlock struct {
	Input1   *synth.TypedInput[int64]
	Input2   *synth.TypedInput[float64]
	Control1 *synth.TypedControl[bool]
	Control2 *synth.TypedControl[string]
	Param1   *synth.TypedParam[int64]
	Param2   *synth.TypedParam[string]
	Output1  *synth.TypedOutput[bool]
	Output2  *synth.TypedOutput[float64]
}

func (tb *TestBlock) Initialize(srate float64) error { return nil }
func (tb *TestBlock) Configure()                     {}
func (tb *TestBlock) Run()                           {}

func TestGetInterface(t *testing.T) {
	tb := &TestBlock{}
	ifc := synth.GetInterface(tb)

	assert.Len(t, ifc.Inputs, 2)
	assert.Len(t, ifc.Controls, 2)
	assert.Len(t, ifc.Params, 2)
	assert.Len(t, ifc.Outputs, 2)

	assert.Contains(t, ifc.Inputs, "Input1")
	assert.Contains(t, ifc.Inputs, "Input2")
	assert.Contains(t, ifc.Controls, "Control1")
	assert.Contains(t, ifc.Controls, "Control2")
	assert.Contains(t, ifc.Params, "Param1")
	assert.Contains(t, ifc.Params, "Param2")
	assert.Contains(t, ifc.Outputs, "Output1")
	assert.Contains(t, ifc.Outputs, "Output2")
}
