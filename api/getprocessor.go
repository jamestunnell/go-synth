package api

import (
	"net/http"

	"github.com/jamestunnell/go-synth/node"
	"github.com/jamestunnell/go-synth/unit/proc/math/abs"
	"github.com/jamestunnell/go-synth/unit/proc/math/add"
	"github.com/jamestunnell/go-synth/unit/proc/math/div"
	"github.com/jamestunnell/go-synth/unit/proc/math/mul"
	"github.com/jamestunnell/go-synth/unit/proc/math/neg"
	"github.com/jamestunnell/go-synth/unit/proc/math/pow"
	"github.com/jamestunnell/go-synth/unit/proc/math/sub"
)

var ProcRegistry *node.CoreRegistry

func init() {
	ProcRegistry = node.NewCoreRegistry()

	ProcRegistry.Register(abs.New())
	ProcRegistry.Register(add.New())
	ProcRegistry.Register(div.New())
	ProcRegistry.Register(mul.New())
	ProcRegistry.Register(neg.New())
	ProcRegistry.Register(pow.New())
	ProcRegistry.Register(sub.New())
}

func getProc(w http.ResponseWriter, r *http.Request) {
	getCore(w, r, ProcRegistry)
}
