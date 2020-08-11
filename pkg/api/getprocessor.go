package api

import (
	"net/http"

	"github.com/jamestunnell/go-synth/pkg/unit/processors"
)

func getProcessor(w http.ResponseWriter, r *http.Request) {
	getUnit(w, r, processors.Builtin)
}
