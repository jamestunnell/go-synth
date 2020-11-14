package api

import (
	"net/http"
)

func getGens(w http.ResponseWriter, r *http.Request) {
	getCores(w, r, GenRegistry)
}
