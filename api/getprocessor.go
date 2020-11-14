package api

import (
	"net/http"

	"github.com/jamestunnell/go-synth/node"
	"github.com/jamestunnell/go-synth/proc/math/abs"
	"github.com/jamestunnell/go-synth/proc/math/add"
	"github.com/jamestunnell/go-synth/proc/math/div"
	"github.com/jamestunnell/go-synth/proc/math/mul"
	"github.com/jamestunnell/go-synth/proc/math/neg"
	"github.com/jamestunnell/go-synth/proc/math/pow"
	"github.com/jamestunnell/go-synth/proc/math/sub"
)

var ProcRegistry *node.Registry

func init() {
	ProcRegistry = node.NewRegistry()

	ProcRegistry.RegisterCore(abs.New())
	ProcRegistry.RegisterCore(add.New())
	ProcRegistry.RegisterCore(div.New())
	ProcRegistry.RegisterCore(mul.New())
	ProcRegistry.RegisterCore(neg.New())
	ProcRegistry.RegisterCore(pow.New())
	ProcRegistry.RegisterCore(sub.New())
}

func getProc(w http.ResponseWriter, r *http.Request) {
	getCore(w, r, ProcRegistry)
}
