package fir_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jamestunnell/go-synth/dsp/filter/fir"
	"github.com/jamestunnell/go-synth/dsp/window"
)

func TestSincFilter(t *testing.T) {
	f, err := fir.NewSincFilter(10000.0, 100.0, 100, window.NewBlackmanHarris())

	if !assert.NoError(t, err) {
		return
	}

	freqContent := f.LowpassResponse()

	_, err := freqContent.MagnitudesDecibel()

	assert.NoError(t, err)

	// t.Log("Lowpass response")
	// for i, freq := range freqContent.Frequencies {
	// 	t.Log(freq, mags[i])
	// }

	// freqContent = f.HighpassResponse()

	// mags, err = freqContent.MagnitudesDecibel()

	// assert.NoError(t, err)

	// t.Log("Highpass response")
	// for i, freq := range freqContent.Frequencies {
	// 	t.Log(freq, mags[i])
	// }
}
