package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

// AddRoutes adds the API routes to the given HTTP router
func AddRoutes(router *mux.Router) {
	router.HandleFunc("/status", getStatus).Methods(http.MethodGet)

	router.HandleFunc("/gens", getGens).Methods(http.MethodGet)
	router.HandleFunc("/gens/{name}", getGen).Methods(http.MethodGet)
	router.HandleFunc("/gens/{name}/render", renderGen).Methods(http.MethodPost)

	router.HandleFunc("/procs", getProcs).Methods(http.MethodGet)
	router.HandleFunc("/procs/{name}", getProc).Methods(http.MethodGet)

	// router.HandleFunc("/nets/create", createNets).Methods(http.MethodPost)
	// router.HandleFunc("/nets/delete", deleteNets).Methods(http.MethodPost)
	// router.HandleFunc("/nets", getAllNets).Methods(http.MethodGet)
	// router.HandleFunc("/nets/{name}", getNetByName).Methods(http.MethodGet)
	// router.HandleFunc("/nets/{name}/render", renderNet).Methods(http.MethodPost)
}
