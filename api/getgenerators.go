package api

import (
	"net/http"

	"github.com/jamestunnell/go-synth/unit/generators"
)

func getGenerators(w http.ResponseWriter, r *http.Request) {
	getUnits(w, r, generators.Builtin)
}
