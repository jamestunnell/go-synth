package api

import (
	"net/http"

	"github.com/jamestunnell/go-synth/pkg/unit/processors"
)

func getProcessors(w http.ResponseWriter, r *http.Request) {
	getUnits(w, r, processors.Builtin)
}
