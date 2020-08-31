package api

import (
	"net/http"

	"github.com/jamestunnell/go-synth/generators"
)

func getGenerator(w http.ResponseWriter, r *http.Request) {
	getUnit(w, r, generators.BuiltinGenerators)
}
