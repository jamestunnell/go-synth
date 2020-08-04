package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/kr/pretty"

	"github.com/jamestunnell/go-synth/pkg/unit"
)

type MakeGeneratorDemoRequest struct {
	Name string `json:"name"`
	SampleRate float64 `json:"srate"`
	Params map[string]float64 `json:"params,omitempty"`
}

const (
	MinDemoSampleRate = 100.0
	MaxDemoSampleRate = 48000.0
)

func makeGeneratorDemo(w http.ResponseWriter, r *http.Request) {
	data, err := ReadRequestData(r)
	if err != nil {
		log.Printf("failed to read request data: %v", err)
		
		w.WriteHeader(http.StatusInternalServerError)
		
		return
	}

	var makeDemoRequest MakeGeneratorDemoRequest
	
	if err = json.Unmarshal(data, &makeDemoRequest); err != nil {
		log.Printf("failed to unmarshal request data: %v", err)
		
		w.WriteHeader(http.StatusBadRequest)
		
		return
	}

	log.Printf("%# v", pretty.Formatter(makeDemoRequest))

	if (makeDemoRequest.SampleRate > MaxDemoSampleRate) || 
		(makeDemoRequest.SampleRate < MinDemoSampleRate) {
		log.Printf("sample rate is not in range [%f,%f]", MinDemoSampleRate, MaxDemoSampleRate)
	
		w.WriteHeader(http.StatusBadRequest)
		
		return
	}

	generatorCore := findGenerator(makeDemoRequest.Name)
	if generatorCore == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	generator := unit.NewUnit(generatorCore)

	var params *unit.Params

	if len(makeDemoRequest.Params) > 0 {
		params = unit.NewParams(makeDemoRequest.Params)
	} else {
		params = unit.NewParams(map[string]float64{})
	}

	generator.FillInDefaultParams(params)

	_, err = generator.Configure(makeDemoRequest.SampleRate, params)
	if err != nil {
		log.Printf("failed to configure generator: %v", err)
		
		w.WriteHeader(http.StatusBadRequest)
		
		return		
	}

	for i := 0; i < 200; i++ {
		outputs := generator.NextSample([]float64{})
		log.Println(outputs)
	}

	w.WriteHeader(http.StatusNoContent)
}