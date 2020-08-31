package api

import (
	"net/http"

	"github.com/jamestunnell/go-synth/generators"
)

func getGenerators(w http.ResponseWriter, r *http.Request) {
	getUnits(w, r, generators.BuiltinGenerators)
}
