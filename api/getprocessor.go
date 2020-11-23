package api

import (
	"net/http"

	"github.com/jamestunnell/go-synth/node"
	"github.com/jamestunnell/go-synth/unit/proc"
)

var ProcRegistry *node.CoreRegistry

func init() {
	ProcRegistry = node.NewCoreRegistry()

	proc.RegisterBuiltin(ProcRegistry)
	proc.RegisterBuiltin(node.WorkingRegistry())
}

func getProc(w http.ResponseWriter, r *http.Request) {
	getCore(w, r, ProcRegistry)
}
