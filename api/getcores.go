package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/jamestunnell/go-synth/node"
)

// CoreInterfacesPayload is returned from the gens/ and procs/ endpoints.
type CoreInterfacesPayload struct {
	Interfaces map[string]*node.Interface `json:"cores"`
}

func getCores(w http.ResponseWriter, r *http.Request, reg *node.CoreRegistry) {
	p := CoreInterfacesPayload{Interfaces: make(map[string]*node.Interface)}

	for _, path := range reg.Paths() {
		core, ok := reg.GetCore(path)
		if !ok {
			log.Printf("failed to make core %s", path)

			w.WriteHeader(http.StatusInternalServerError)

			return
		}

		p.Interfaces[path] = core.Interface()
	}

	data, err := json.Marshal(p)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(data)
}
