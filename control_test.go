package synth_test

import (
	"testing"

	"github.com/jamestunnell/go-synth"
)

func TestNewControls(t *testing.T) {
	testNewInput(t, synth.NewUint64Control(0), "uint64")
	testNewInput(t, synth.NewInt64Control(0), "int64")
	testNewInput(t, synth.NewFloat64Control(0), "float64")
	testNewInput(t, synth.NewBoolControl(false), "bool")
	testNewInput(t, synth.NewStringControl(""), "string")
}
