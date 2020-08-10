package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

func AddRoutes(router *mux.Router) {
	router.HandleFunc("/status", getStatus).Methods(http.MethodGet)
	router.HandleFunc("/unit/generators", getGenerators).Methods(http.MethodGet)
	router.HandleFunc("/unit/generators/{name}/{srate}", getGenerator).Methods(http.MethodGet)
	router.HandleFunc("/unit/generators/{name}/{srate}/demo", makeGeneratorDemo).Methods(http.MethodPost)
}
