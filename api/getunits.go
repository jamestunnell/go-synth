package api

import (
	"encoding/json"
	"net/http"

	"github.com/jamestunnell/go-synth/unit"
)

type GetUnitsPayload struct {
	Units []*unit.BasicInfo `json:"units"`
}

func getUnits(w http.ResponseWriter, r *http.Request, plugins []*unit.Plugin) {
	info := make([]*unit.BasicInfo, len(plugins))

	for i, plugin := range plugins {
		info[i] = plugin.BasicInfo
	}

	p := GetUnitsPayload{Units: info}

	data, err := json.Marshal(p)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(data)
}
