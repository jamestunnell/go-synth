package api

import (
	"net/http"

	"github.com/jamestunnell/go-synth/processors"
)

func getProcessor(w http.ResponseWriter, r *http.Request) {
	getUnit(w, r, processors.BuiltinProcessors)
}
