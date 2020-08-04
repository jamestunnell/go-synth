package api

import (
	"encoding/json"
	"net/http"

	"github.com/jamestunnell/go-synth/pkg/metadata"
	"github.com/jamestunnell/go-synth/pkg/unit/generators"
)

var (
	builtinGenerators = []*metadata.Metadata{
		generators.SquareWaveMetadata,
	}
)

type GetGeneratorsPayload struct {
	Generators []*metadata.Metadata `json:"generators"`
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
