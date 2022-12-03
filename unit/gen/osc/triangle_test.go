package osc_test

import (
	"testing"

	"github.com/jamestunnell/go-synth"
	"github.com/jamestunnell/go-synth/unit/gen/osc"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTriangle(t *testing.T) {
	const (
		srate    = 15.0
		oscFreq  = 3.0
		oscPhase = 0.0
	)

	freq := synth.NewConstBlock[float64](oscFreq)
	phase := synth.NewConstBlock[float64](oscPhase)
	osc := osc.NewTriangle()

	require.NoError(t, osc.Freq.Connect(freq.Out))
	require.NoError(t, osc.Phase.Connect(phase.Out))

	require.NoError(t, freq.Initialize(srate, 1))
	require.NoError(t, phase.Initialize(srate, 1))
	require.NoError(t, osc.Initialize(srate, 15))

	osc.Configure()
	osc.Run()

	outBuf := osc.Out.Buffer().([]float64)

	// First 5 samples should contain a complete cycle
	assert.InDelta(t, -1.0, outBuf[0], 1e-5)
	assert.InDelta(t, -0.2, outBuf[1], 1e-5)
	assert.InDelta(t, 0.6, outBuf[2], 1e-5)
	assert.InDelta(t, 0.6, outBuf[3], 1e-5)
	assert.InDelta(t, -0.2, outBuf[4], 1e-5)

	// Then the first cycle should be repeated twice
	assert.Equal(t, outBuf[:5], outBuf[5:10])
	assert.Equal(t, outBuf[:5], outBuf[10:15])
}
