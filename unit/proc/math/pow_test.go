package math_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jamestunnell/go-synth"
	"github.com/jamestunnell/go-synth/unit/gen/array"
	"github.com/jamestunnell/go-synth/unit/proc/math"
)

var inVals = []float64{1.0, 0.5, -0.5}

func TestPowInvert(t *testing.T) {
	testPow(t, -1, []float64{1.0, 0.5, -0.5}, []float64{1.0, 2.0, -2.0})
}

func TestPowOnes(t *testing.T) {
	testPow(t, 0, []float64{1.0, 0.5, -0.5}, []float64{1.0, 1.0, 1.0})
}

func TestPowIdentity(t *testing.T) {
	testPow(t, 1, []float64{1.0, 0.5, -0.5}, []float64{1.0, 0.5, -0.5})
}

func TestPowSquare(t *testing.T) {
	testPow(t, 2, []float64{1.0, 0.5, -0.5}, []float64{1.0, 0.25, 0.25})
}

func TestPowCube(t *testing.T) {
	testPow(t, 3, []float64{1.0, 0.5, -0.5}, []float64{1.0, 0.125, -0.125})
}

func testPow(t *testing.T, exp float64, inVals, outVals []float64) {
	in := array.NewOneshot(inVals...)
	blk := math.NewPow()
	expCtrl := synth.NewConst(exp)

	assert.NoError(t, blk.In.Connect(in.Out))
	assert.NoError(t, blk.Exp.Connect(expCtrl.Out))

	assert.NoError(t, expCtrl.Initialize(100.0/3, 1))
	assert.NoError(t, in.Initialize(100.0, len(inVals)))
	assert.NoError(t, blk.Initialize(100.0, len(inVals)))

	blk.Configure()

	expCtrl.Run()
	in.Run()
	blk.Run()

	for i, outVal := range outVals {
		assert.Equal(t, outVal, blk.Out.Buffer[i])
	}
}
