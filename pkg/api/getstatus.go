package api

import (
	"net/http"
)

func getStatus(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNoContent)
}
