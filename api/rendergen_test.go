package api_test

import (
	"net/http"
	"os"
	"testing"

	"github.com/jamestunnell/go-synth/api"
	"github.com/jamestunnell/go-synth/node"
	"github.com/jamestunnell/go-synth/unit/gen/osc"
	"github.com/stretchr/testify/assert"
)

func TestProcessRenderRequestBadSrate(t *testing.T) {
	for _, srate := range []float64{0.0, -100.0, 20000.0} {
		req := &api.RenderGenRequest{
			DurSec:     1.0,
			BitDepth:   16,
			SampleRate: srate,
			Controls:   map[string]float64{"Freq": 220.0},
			Params:     node.ParamMap{"Wave": 0},
		}
		core := &osc.Osc{}

		testProcessRenderRequestBadRequest(t, req, core)
	}
}

func TestProcessRenderRequestBadBitDepth(t *testing.T) {
	for _, bdepth := range []int{0, -1} {
		req := &api.RenderGenRequest{
			DurSec:     1.0,
			BitDepth:   bdepth,
			SampleRate: 22050,
			Controls:   map[string]float64{"Freq": 220.0},
			Params:     node.ParamMap{"Wave": 0},
		}
		core := &osc.Osc{}

		testProcessRenderRequestBadRequest(t, req, core)
	}
}

func TestProcessRenderRequestBadDur(t *testing.T) {
	for _, dur := range []float64{0.0, -0.1, api.MaxDurSec + 0.1} {
		req := &api.RenderGenRequest{
			DurSec:     dur,
			BitDepth:   16,
			SampleRate: 22050,
			Controls:   map[string]float64{"Freq": 220.0},
			Params:     node.ParamMap{"Wave": 0},
		}
		core := &osc.Osc{}

		testProcessRenderRequestBadRequest(t, req, core)
	}
}

func TestProcessRenderRequestMissingParam(t *testing.T) {
	req := &api.RenderGenRequest{
		DurSec:     1.0,
		BitDepth:   16,
		SampleRate: 22050,
		Controls:   map[string]float64{"Freq": 220.0},
		Params:     node.ParamMap{},
	}
	core := &osc.Osc{}

	testProcessRenderRequestBadRequest(t, req, core)
}

func testProcessRenderRequestBadRequest(t *testing.T, req *api.RenderGenRequest, core node.Core) {
	wavFile, e := api.ProcessRenderRequest(req, core)

	assert.Error(t, e)
	assert.Equal(t, http.StatusBadRequest, e.StatusCode())
	assert.Nil(t, wavFile)
}

func TestProcessRenderRequestHappyPath(t *testing.T) {
	for i := 0; i < 4; i++ {
		req := &api.RenderGenRequest{
			BitDepth:   16,
			DurSec:     1.0,
			SampleRate: 22050,
			Controls:   map[string]float64{"Freq": 220.0},
			Params:     node.ParamMap{"Wave": float64(i)},
		}
		core := &osc.Osc{}
		wavFile, e := api.ProcessRenderRequest(req, core)

		assert.NoError(t, e)
		assert.NotNil(t, wavFile)

		defer os.Remove(wavFile.Name())

		// t.Logf("wave type=%d, file name=%s", i, wavFile.Name())
	}
}

func TestProcessRenderRequestNumSamplesNotMultipleOfChunkSize(t *testing.T) {
	req := &api.RenderGenRequest{
		BitDepth:   16,
		DurSec:     1.2,
		SampleRate: 44100,
		Controls:   map[string]float64{"Freq": 440.0},
		Params:     node.ParamMap{"Wave": float64(0)},
	}
	core := &osc.Osc{}
	wavFile, e := api.ProcessRenderRequest(req, core)

	assert.NoError(t, e)
	assert.NotNil(t, wavFile)

	defer os.Remove(wavFile.Name())
}
