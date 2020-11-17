package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jamestunnell/go-synth/node"
)

func getCore(w http.ResponseWriter, r *http.Request, reg *node.CoreRegistry) {
	vars := mux.Vars(r)

	core := findCore(vars["name"], reg)
	if core == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	data, err := json.Marshal(core.Interface())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(data)
}
