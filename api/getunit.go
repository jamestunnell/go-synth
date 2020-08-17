package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jamestunnell/go-synth/unit"
)

func getUnit(w http.ResponseWriter, r *http.Request, plugins []*unit.Plugin) {
	vars := mux.Vars(r)

	plugin := findPlugin(vars["name"], plugins)
	if plugin == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	data, err := json.Marshal(plugin)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(data)
}
