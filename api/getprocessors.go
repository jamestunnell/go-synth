package api

import (
	"net/http"

	"github.com/jamestunnell/go-synth/processors"
)

func getProcessors(w http.ResponseWriter, r *http.Request) {
	getUnits(w, r, processors.BuiltinProcessors)
}
