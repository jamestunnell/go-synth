package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func getGenerator(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	plugin := findGeneratorPlugin(vars["name"])
	if plugin == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	srate, ok := getSrate(vars)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	info := &GeneratorInfo{
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
