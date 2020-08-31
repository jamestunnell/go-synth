package api

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"reflect"

	"github.com/gorilla/mux"
	"github.com/kr/pretty"

	"github.com/jamestunnell/go-synth/generators"
	"github.com/jamestunnell/go-synth/node"
	"github.com/jamestunnell/go-synth/util"
)

type MakeGeneratorDemoRequest struct {
	DurSec     float64            `json:"dursec"`
	SampleRate float64            `json:"srate"`
	Params     map[string]float64 `json:"params,omitempty"`
}

const (
	DefaultDurSec = MinDurSec
	DemoBitDepth  = 16
	ChunkSize     = 50
	MinDurSec     = 1.0
	MaxDurSec     = 10.0
)

func makeGeneratorDemo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	core := findUnit(vars["name"], generators.BuiltinGenerators)
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

	genNode, err := createGenNode(core, request.Params)
	if err != nil {
		log.Printf("failed to create gen node: %v", err)

		w.WriteHeader(http.StatusBadRequest)

		return
	}

	genNode.Initialize(request.SampleRate)

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

func createGenNode(core node.Core, requestParams map[string]float64) (*node.Node, error) {
	// clone the built-in core
	vCore := reflect.New(reflect.ValueOf(core).Elem().Type())
	ifc := core.GetInterface()

	// get the parameter values from the request
	for paramName := range ifc.Parameters {
		val, found := requestParams[paramName]
		if found {
			vCore.Elem().FieldByName(paramName).Set(reflect.ValueOf(val))
		}
	}

	core2 := vCore.Interface().(node.Core)

	return node.MakeNode(core2, ChunkSize)
}
