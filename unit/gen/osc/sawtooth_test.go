package osc_test

import (
	"testing"

	"github.com/jamestunnell/go-synth"
	"github.com/jamestunnell/go-synth/unit/gen/osc"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSawtooth(t *testing.T) {
	const (
		srate    = 15.0
		oscFreq  = 3.0
		oscPhase = 0.0
	)

	freq := synth.NewConst(oscFreq)
	phase := synth.NewConst(oscPhase)
	osc := osc.NewSawtooth()

	require.NoError(t, osc.Freq.Connect(freq.Out))
	require.NoError(t, osc.Phase.Connect(phase.Out))

	require.NoError(t, freq.Initialize(srate, 1))
	require.NoError(t, phase.Initialize(srate, 1))
	require.NoError(t, osc.Initialize(srate, 15))

	osc.Configure()
	osc.Run()

	outBuf := osc.Out.Buffer

	// First 5 samples should contain a complete cycle
	assert.Equal(t, 0.0, outBuf[0])
	assert.Equal(t, 0.4, outBuf[1])
	assert.Equal(t, 0.8, outBuf[2])
	assert.Equal(t, -0.8, outBuf[3])
	assert.Equal(t, -0.4, outBuf[4])

	// Then the first cycle should be repeated twice
	assert.Equal(t, outBuf[:5], outBuf[5:10])
	assert.Equal(t, outBuf[:5], outBuf[10:15])
}
