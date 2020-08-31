package api

import (
	"encoding/json"
	"net/http"
	"reflect"

	"github.com/jamestunnell/go-synth/node"
)

type GetUnitsPayload struct {
	Units map[string]*node.Interface `json:"units"`
}

func getUnits(w http.ResponseWriter, r *http.Request, cores []node.Core) {
	p := GetUnitsPayload{Units: make(map[string]*node.Interface)}

	for _, core := range cores {
		p.Units[reflect.TypeOf(core).Elem().Name()] = core.GetInterface()
	}

	data, err := json.Marshal(p)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(data)
}
