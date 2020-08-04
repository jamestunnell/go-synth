package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func getGenerator(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	generator := findGenerator(vars["name"])
	if generator == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	data, err := json.Marshal(generator.Metadata())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(data)
}
