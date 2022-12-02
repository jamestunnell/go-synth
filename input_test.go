package synth_test

import (
	"testing"

	"github.com/jamestunnell/go-synth"
	"github.com/stretchr/testify/assert"
)

func TestNewInputs(t *testing.T) {
	testNewInput(t, synth.NewUint64Input(), "uint64")
	testNewInput(t, synth.NewInt64Input(), "int64")
	testNewInput(t, synth.NewFloat64Input(), "float64")
	testNewInput(t, synth.NewBoolInput(), "bool")
	testNewInput(t, synth.NewStringInput(), "string")
}

func testNewInput(t *testing.T, in synth.Input, expectedType string) {
	assert.Equal(t, expectedType, in.Type())
	assert.False(t, in.Connected())
}
