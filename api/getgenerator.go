package api

import (
	"net/http"

	"github.com/jamestunnell/go-synth/node"
	"github.com/jamestunnell/go-synth/unit/gen"
)

// GenRegistry holds just the registered unit generators
var GenRegistry *node.CoreRegistry

func init() {
	GenRegistry = node.NewCoreRegistry()

	gen.RegisterBuiltin(GenRegistry)
	gen.RegisterBuiltin(node.WorkingRegistry())
}

func getGen(w http.ResponseWriter, r *http.Request) {
	getCore(w, r, GenRegistry)
}
