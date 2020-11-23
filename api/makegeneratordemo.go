package api

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/kr/pretty"

	"github.com/jamestunnell/go-synth/node"
	"github.com/jamestunnell/go-synth/util"
)

type MakeGeneratorDemoRequest struct {
	DurSec     float64                `json:"dursec"`
	SampleRate float64                `json:"srate"`
	Controls   map[string]float64     `json:"controls,omitempty"`
	Params     map[string]interface{} `json:"params,omitempty"`
}

const (
	DefaultDurSec = MinDurSec
	DemoBitDepth  = 16
	ChunkSize     = 50
	MinDurSec     = 1.0
	MaxDurSec     = 10.0
)

func renderGen(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	core := findCore(vars["name"], GenRegistry)
	if core == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	data, err := ReadRequestData(r)
	if err != nil {
		log.Printf("failed to read request data: %v", err)

		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	var request MakeGeneratorDemoRequest

	if err = json.Unmarshal(data, &request); err != nil {
		log.Printf("failed to unmarshal request JSON object: %v", err)

		w.WriteHeader(http.StatusBadRequest)

		return
	}

	log.Printf("make demo request: %# v", pretty.Formatter(request))

	durSec := request.DurSec
	if durSec == 0.0 {
		durSec = DefaultDurSec
	}

	if durSec > MaxDurSec || durSec < MinDurSec {
		log.Printf("duration %f is not in range [%f,%f]", durSec, MinDurSec, MaxDurSec)

		w.WriteHeader(http.StatusBadRequest)

		return
	}

	if !isSampleRateValid(request.SampleRate) {
		log.Printf("sample rate %f is invalid", request.SampleRate)

		w.WriteHeader(http.StatusBadRequest)

		return
	}

	genNode := createGenNode(core, request.Params, request.Controls)
	// if err != nil {
	// 	log.Printf("failed to create gen node: %v", err)

	// 	w.WriteHeader(http.StatusBadRequest)

	// 	return
	// }

	genNode.Initialize(request.SampleRate, ChunkSize)

	renderParams := &util.RenderParams{
		DurSec:     durSec,
		BitDepth:   DemoBitDepth,
		SampleRate: int(request.SampleRate),
	}

	// We can use a pattern of "pre-*.txt" to get an extension like: /tmp/pre-123456.txt
	tmpFile, err := ioutil.TempFile(os.TempDir(), "demo-*.wav")
	if err != nil {
		log.Printf("failed to create temporary file: %v", err)

		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	defer os.Remove(tmpFile.Name())

	err = util.RenderWAV(genNode, tmpFile, renderParams)
	if err != nil {
		log.Printf("failed to render WAV: %v", err)

		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	tmpFile.Close()

	wavFileName := tmpFile.Name()

	w.Header().Set("Content-Type", "audio/wav")
	http.ServeFile(w, r, wavFileName)
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
