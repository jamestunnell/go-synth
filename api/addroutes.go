package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

func AddRoutes(router *mux.Router) {
	router.HandleFunc("/status", getStatus).Methods(http.MethodGet)
	router.HandleFunc("/unit/generators", getGenerators).Methods(http.MethodGet)
	router.HandleFunc("/unit/processors", getProcessors).Methods(http.MethodGet)
	router.HandleFunc("/unit/generators/{name}", getGenerator).Methods(http.MethodGet)
	router.HandleFunc("/unit/processors/{name}", getProcessor).Methods(http.MethodGet)
	router.HandleFunc("/unit/generators/{name}/demo", makeGeneratorDemo).Methods(http.MethodPost)
}
