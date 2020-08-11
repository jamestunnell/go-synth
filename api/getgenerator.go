package api

import (
	"net/http"

	"github.com/jamestunnell/go-synth/unit/generators"
)

func getGenerator(w http.ResponseWriter, r *http.Request) {
	getUnit(w, r, generators.Builtin)
}
