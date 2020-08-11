package api

import (
	"net/http"

	"github.com/jamestunnell/go-synth/pkg/unit/generators"
)

func getGenerators(w http.ResponseWriter, r *http.Request) {
	getUnits(w, r, generators.Builtin)
}
