package osc_test

import (
	"testing"

	"github.com/jamestunnell/go-synth"
	"github.com/jamestunnell/go-synth/unit/gen/osc"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSquare(t *testing.T) {
	const (
		srate    = 15.0
		oscFreq  = 3.0
		oscPhase = 0.0
	)

	freq := synth.NewConst[float64](oscFreq)
	phase := synth.NewConst[float64](oscPhase)
	osc := osc.NewSquare()

	require.NoError(t, osc.Freq.Connect(freq.Out))
	require.NoError(t, osc.Phase.Connect(phase.Out))

	require.NoError(t, freq.Initialize(srate, 1))
	require.NoError(t, phase.Initialize(srate, 1))
	require.NoError(t, osc.Initialize(srate, 15))

	osc.Configure()
	osc.Run()

	outBuf := osc.Out.BufferValues

	// First 5 samples should contain a complete cycle
	assert.Equal(t, 1.0, outBuf[0])
	assert.InDelta(t, 1.0, outBuf[1], 1e-3)
	assert.InDelta(t, 1.0, outBuf[2], 1e-3)
	assert.InDelta(t, -1.0, outBuf[3], 1e-3)
	assert.InDelta(t, -1.0, outBuf[4], 1e-3)

	// Then the first cycle should be repeated twice
	assert.Equal(t, outBuf[:5], outBuf[5:10])
	assert.Equal(t, outBuf[:5], outBuf[10:15])
}
