package api

import (
	"net/http"

	"github.com/jamestunnell/go-synth/unit/processors"
)

func getProcessor(w http.ResponseWriter, r *http.Request) {
	getUnit(w, r, processors.Builtin)
}
