package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/kr/pretty"

	"github.com/jamestunnell/go-synth/node"
	"github.com/jamestunnell/go-synth/util"
	"github.com/jamestunnell/go-synth/util/httperr"
)

type RenderGenRequest struct {
	DurSec     float64                `json:"dursec"`
	SampleRate float64                `json:"srate"`
	BitDepth   int                    `json:"bitdepth"`
	Controls   map[string]float64     `json:"controls,omitempty"`
	Params     map[string]interface{} `json:"params,omitempty"`
}

const (
	ChunkSize = 50
	MaxDurSec = 1000.0
)

func renderGen(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]

	core := findCore(name, GenRegistry)
	if core == nil {
		msg := fmt.Sprintf("failed to find core with %s in path", name)
		e := httperr.New404(msg)

		e.WriteError(w)

		return
	}

	data, err := ReadRequestData(r)
	if err != nil {
		msg := fmt.Sprintf("failed to read request data: %v", err)
		e := httperr.New500(msg)

		e.WriteError(w)

		return
	}

	var request RenderGenRequest

	if err = json.Unmarshal(data, &request); err != nil {
		msg := fmt.Sprintf("failed to unmarshal request JSON object: %v", err)
		e := httperr.New400(msg)

		e.WriteError(w)

		return
	}

	log.Printf("make demo request: %# v", pretty.Formatter(request))

	wavFile, e := ProcessRenderRequest(&request, core)
	if e != nil {
		e.WriteError(w)
		return
	}

	log.Printf("rendered %s", wavFile.Name())

	w.Header().Set("Content-Type", "audio/wav")
	http.ServeFile(w, r, wavFile.Name())
}

func ProcessRenderRequest(request *RenderGenRequest, core node.Core) (*os.File, httperr.E) {
	if request.BitDepth <= 0 {
		msg := fmt.Sprintf("bit depth %d is not positive", request.BitDepth)
		return nil, httperr.New400(msg)
	}

	if request.DurSec > MaxDurSec || request.DurSec <= 0.0 {
		msg := fmt.Sprintf("duration %f is not in range (%f,%f]", request.DurSec, 0.0, MaxDurSec)
		return nil, httperr.New400(msg)
	}

	if !isSampleRateValid(request.SampleRate) {
		msg := fmt.Sprintf("sample rate %f is invalid", request.SampleRate)
		return nil, httperr.New400(msg)
	}

	genNode := createGenNode(core, request.Params, request.Controls)

	err := genNode.Initialize(request.SampleRate, ChunkSize)
	if err != nil {
		msg := fmt.Sprintf("failed to initialize gen node: %v", err)
		return nil, httperr.New400(msg)
	}

	renderParams := &util.RenderParams{
		DurSec:     request.DurSec,
		BitDepth:   request.BitDepth,
		SampleRate: int(request.SampleRate),
	}

	// We can use a pattern of "pre-*.txt" to get an extension like: /tmp/pre-123456.txt
	tmpFile, err := ioutil.TempFile(os.TempDir(), "demo-*.wav")
	if err != nil {
		msg := fmt.Sprintf("failed to create temporary file: %v", err)
		return nil, httperr.New500(msg)
	}

	err = util.RenderWAV(genNode, tmpFile, renderParams)
	if err != nil {
		msg := fmt.Sprintf("failed to render WAV: %v", err)
		return nil, httperr.New500(msg)
	}

	tmpFile.Close()

	return tmpFile, nil
}

func isSampleRateValid(srate float64) bool {
	switch srate {
	case 22050:
		return true
	case 44100:
		return true
	case 48000:
		return true
	case 96000:
		return true
	case 192000:
		return true
	}

	return false
}

func createGenNode(core node.Core, params node.ParamMap, controlVals map[string]float64) *node.Node {
	controls := node.Map{}
	for name, val := range controlVals {
		controls[name] = node.NewConst(val)
	}

	return &node.Node{
		Core:     core,
		Inputs:   node.Map{},
		Controls: controls,
		Params:   params,
	}
}
