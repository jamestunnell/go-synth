package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jamestunnell/go-synth/node"
)

func getUnit(w http.ResponseWriter, r *http.Request, cores []node.Core) {
	vars := mux.Vars(r)

	core := findUnit(vars["name"], cores)
	if core == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	data, err := json.Marshal(core.GetInterface())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(data)
}
