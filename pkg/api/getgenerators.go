package api

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

var (
	builtinGenerators = []*GeneratorInfo{
		{
			ID:          uuid.MustParse("9c575a62-8d7a-46c5-ab72-92d342f3e238"),
			Name:        "Square wave oscillator",
			Author:      "James Tunnell",
			Description: "50% duty cycle square wave oscillates from -1 to 1",
		},
	}
)

type GetGeneratorsPayload struct {
	Generators []*GeneratorInfo `json:"generators"`
}

func getGenerators(w http.ResponseWriter, r *http.Request) {
	p := GetGeneratorsPayload{Generators: builtinGenerators}

	data, err := json.Marshal(p)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(data)
}
