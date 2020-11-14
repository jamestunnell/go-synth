package api

import (
	"net/http"

	"github.com/jamestunnell/go-synth/gen/array/oneshot"
	"github.com/jamestunnell/go-synth/gen/array/repeat"
	"github.com/jamestunnell/go-synth/gen/osc/saw"
	"github.com/jamestunnell/go-synth/gen/osc/sine"
	"github.com/jamestunnell/go-synth/gen/osc/square"
	"github.com/jamestunnell/go-synth/gen/osc/triangle"
	"github.com/jamestunnell/go-synth/node"
)

var GenRegistry *node.Registry

func init() {
	GenRegistry = node.NewRegistry()

	GenRegistry.RegisterCore(oneshot.New([]float64{}))
	GenRegistry.RegisterCore(repeat.New([]float64{}))
	GenRegistry.RegisterCore(saw.New())
	GenRegistry.RegisterCore(sine.New())
	GenRegistry.RegisterCore(square.New())
	GenRegistry.RegisterCore(triangle.New())
}

func getGen(w http.ResponseWriter, r *http.Request) {
	getCore(w, r, GenRegistry)
}
