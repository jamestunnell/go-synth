package api

import (
	"net/http"
)

func getProcs(w http.ResponseWriter, r *http.Request) {
	getCores(w, r, ProcRegistry)
}
