package api

import (
	"encoding/json"
	"net/http"

	"github.com/jamestunnell/go-synth/pkg/unit"
	"github.com/jamestunnell/go-synth/pkg/unit/generators"
)

type GetGeneratorsPayload struct {
	Generators []*unit.BasicInfo `json:"generators"`
}

func getGenerators(w http.ResponseWriter, r *http.Request) {
	info := make([]*unit.BasicInfo, len(generators.Builtin))

	for i, plugin := range generators.Builtin {
		info[i] = plugin.BasicInfo
	}

	p := GetGeneratorsPayload{Generators: info}

	data, err := json.Marshal(p)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(data)
}
