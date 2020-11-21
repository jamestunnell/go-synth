package api

import (
	"net/http"

	"github.com/jamestunnell/go-synth/unit/gen/array/oneshot"
	"github.com/jamestunnell/go-synth/unit/gen/array/repeat"
	"github.com/jamestunnell/go-synth/unit/gen/osc/saw"
	"github.com/jamestunnell/go-synth/unit/gen/osc/sine"
	"github.com/jamestunnell/go-synth/unit/gen/osc/square"
	"github.com/jamestunnell/go-synth/unit/gen/osc/triangle"
	"github.com/jamestunnell/go-synth/node"
)

var GenRegistry *node.CoreRegistry

func init() {
	GenRegistry = node.NewCoreRegistry()

	GenRegistry.Register(oneshot.New([]float64{}))
	GenRegistry.Register(repeat.New([]float64{}))
	GenRegistry.Register(saw.New())
	GenRegistry.Register(sine.New())
	GenRegistry.Register(square.New())
	GenRegistry.Register(triangle.New())
}

func getGen(w http.ResponseWriter, r *http.Request) {
	getCore(w, r, GenRegistry)
}
