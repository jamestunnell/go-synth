package node

import (
	"encoding/json"
	"fmt"

	"github.com/jamestunnell/go-synth/util/param"
)

// Map maps names to nodes
type Map map[string]*Node

// Mod function is used to alter a node on creation,
// e.g. by adding an input, control, or param
type Mod func(*Node)

// Node provides the framework for running a Core.
type Node struct {
	// Inputs provide functional input to the core every sample.
	Inputs Map
	// Controls provide configuration input to the core once every
	// chunksize samples.
	Controls Map
	// Params provide configuration input to the core once at init time
	Params param.Map
	// Mul will, if not nil, be multiplied with the node output (before any addition)
	Mul *Node
	// Add will, if not nil, be added to the node output (after any multiplication)
	Add *Node

	core        Core
	corePath    string
	output      *Buffer
	initialized bool
}

// nodeStore is used to serialize/deserialize node as JSON
type nodeStore struct {
	Inputs   Map       `json:"inputs,omitempty"`
	Controls Map       `json:"controls,omitempty"`
	Params   param.Map `json:"params,omitempty"`
	CorePath string    `json:"corePath"`
	Mul      *Node     `json:"mul,omitempty"`
	Add      *Node     `json:"add,omitempty"`
}

// New makes a new node and applies the given mods.
func New(c Core, mods ...Mod) *Node {
	n := &Node{
		core:     c,
		corePath: CorePath(c),
		Inputs:   Map{},
		Controls: Map{},
		Params:   param.Map{},
	}

	for _, mod := range mods {
		mod(n)
	}

	return n
}

// Core returns the node core.
func (n *Node) Core() Core {
	return n.core
}

// CorePath returns the path of node core underlying type.
func (n *Node) CorePath() string {
	return n.corePath
}

// Output returns the latest results from running the core.
// Returns nil before the node is initialized.
func (n *Node) Output() *Buffer {
	return n.output
}

// Initialized returns true if the node has been initialized.
func (n *Node) Initialized() bool {
	return n.initialized
}

// Initialize initializes the node and its dependencies.
// Dependencies are initialized first.
func (n *Node) Initialize(srate float64, depth int) error {
	ifc := n.core.Interface()

	ifc.EnsureControls(n.Controls)

	if err := n.Validate(ifc); err != nil {
		return err
	}

	// Initialize inputs
	for _, inputNode := range n.Inputs {
		if err := inputNode.Initialize(srate, depth); err != nil {
			return err
		}
	}

	// Initialize controls
	controlSampleRate := srate / float64(depth)
	for _, controlNode := range n.Controls {
		if err := controlNode.Initialize(controlSampleRate, 1); err != nil {
			return err
		}
	}

	args := &InitArgs{
		SampleRate: srate,
		Inputs:     n.Inputs,
		Controls:   n.Controls,
		Params:     n.Params,
	}

	// Initialize the core
	if err := n.core.Initialize(args); err != nil {
		return fmt.Errorf("failed to initialize core: %w", err)
	}

	// Initialize Mul and Add
	if err := n.initMulAdd(srate, depth); err != nil {
		return err
	}

	n.output = NewBuffer(depth)
	n.initialized = true

	return nil
}

// Run runs the node and its dependencies. Dependencies are
// ran first. The node core is configured before running.
func (n *Node) Run() {
	if !n.initialized {
		panic("node is not initialized")
	}

	for _, inputNode := range n.Inputs {
		inputNode.Run()
	}

	for _, controlNode := range n.Controls {
		controlNode.Run()
	}

	n.core.Configure()

	n.core.Run(n.output)

	n.runMulAdd()
}

// MarshalJSON generates node JSON data.
// Returns a non-nil error in case of failure.
func (n *Node) MarshalJSON() ([]byte, error) {
	// Use an anonymous struct which has public fields
	// (needed for json.Marshal to work)
	ns := nodeStore{
		Inputs:   n.Inputs,
		Controls: n.Controls,
		Params:   n.Params,
		CorePath: n.corePath,
		Mul:      n.Mul,
		Add:      n.Add,
	}

	return json.Marshal(ns)
}

// UnmarshalJSON restores a node from the given JSON data.
// Returns a non-nil error in case of failure.
func (n *Node) UnmarshalJSON(data []byte) error {
	var ns nodeStore

	err := json.Unmarshal(data, &ns)
	if err != nil {
		return err
	}

	// find the core using the path
	core, ok := WorkingRegistry().GetCore(ns.CorePath)
	if !ok {
		return fmt.Errorf("failed to find core path %s in working registry", ns.CorePath)
	}

	n.Inputs = ns.Inputs
	n.Controls = ns.Controls
	n.Params = ns.Params
	n.core = core
	n.corePath = ns.CorePath
	n.Mul = ns.Mul
	n.Add = ns.Add

	if n.Inputs == nil {
		n.Inputs = Map{}
	}

	if n.Controls == nil {
		n.Controls = Map{}
	}

	if n.Params == nil {
		n.Params = param.Map{}
	}

	return nil
}

// Validate checks that the node has all of the inputs and params
// in the core interface.
func (n *Node) Validate(ifc *Interface) error {
	if err := ifc.CheckInputs(n.Inputs); err != nil {
		return err
	}

	if err := ifc.CheckParams(n.Params); err != nil {
		return err
	}

	return nil
}

func (n *Node) initMulAdd(srate float64, depth int) error {
	if n.Mul != nil {
		if err := n.Mul.Initialize(srate, depth); err != nil {
			return err
		}
	}

	if n.Add != nil {
		if err := n.Add.Initialize(srate, depth); err != nil {
			return err
		}
	}

	return nil
}

func (n *Node) runMulAdd() {
	if n.Mul != nil {
		n.Mul.Run()

		mulBuf := n.Mul.Output()

		for i := 0; i < n.output.Length; i++ {
			n.output.Values[i] *= mulBuf.Values[i]
		}
	}

	if n.Add != nil {
		n.Add.Run()

		addBuf := n.Add.Output()

		for i := 0; i < n.output.Length; i++ {
			n.output.Values[i] += addBuf.Values[i]
		}
	}
}
