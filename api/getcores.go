package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/jamestunnell/go-synth/node"
)

type GetCoresPayload struct {
	Cores map[string]*node.Interface `json:"cores"`
}

func getCores(w http.ResponseWriter, r *http.Request, reg *node.Registry) {
	p := GetCoresPayload{Cores: make(map[string]*node.Interface)}

	for _, path := range reg.Paths() {
		core, ok := reg.MakeCore(path)
		if !ok {
			log.Printf("failed to make core %s", path)

			w.WriteHeader(http.StatusInternalServerError)

			return
		}

		p.Cores[path] = core.Interface()
	}

	data, err := json.Marshal(p)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(data)
}
