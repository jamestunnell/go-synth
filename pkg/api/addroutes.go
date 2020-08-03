package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

func AddRoutes(router *mux.Router) {
	router.HandleFunc("/status", getStatus).Methods(http.MethodGet)
	router.HandleFunc("/generators", getGenerators).Methods(http.MethodGet)
}
