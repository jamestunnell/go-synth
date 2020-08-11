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

	srate, ok := getSrate(vars)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	info := &UnitInfo{
		BasicInfo: plugin.BasicInfo,
		Interface: plugin.GetInterface(float64(srate)),
		ExtraInfo: plugin.ExtraInfo,
	}

	data, err := json.Marshal(info)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(data)
}
