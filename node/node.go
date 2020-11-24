package node

import (
	"encoding/json"
	"fmt"

	"github.com/buger/jsonparser"

	"github.com/jamestunnell/go-synth/util/param"
)

// Map is a type alias
type Map = map[string]*Node

type Mod func(*Node)

// Node provides the framework for running a Core.
type Node struct {
	// Inputs provide functional input to the core every sample.
	Inputs Map `json:"inputs"`
	// Controls provide configuration input to the core once every
	// chunksize samples.
	Controls Map `json:"controls"`
	// Params provide configuration input to the core once at init time
	Params      param.Map `json:"params"`
	core        Core
	corePath    string
	output      *Buffer
	initialized bool
}

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

	if err := n.validate(ifc); err != nil {
		return err
	}

	for _, inputNode := range n.Inputs {
		if err := inputNode.Initialize(srate, depth); err != nil {
			return err
		}
	}

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

	if err := n.core.Initialize(args); err != nil {
		return fmt.Errorf("failed to initialize core: %w", err)
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
}

// MarshalJSON generates node JSON data.
// Returns a non-nil error in case of failure.
func (n *Node) MarshalJSON() ([]byte, error) {
	// Use an anonymous struct which has public fields
	// (needed for json.Marshal to work)
	n2 := struct {
		Inputs   Map       `json:"inputs"`
		Controls Map       `json:"controls"`
		Params   param.Map `json:"params"`
		CorePath string    `json:"corePath"`
	}{
		Inputs:   n.Inputs,
		Controls: n.Controls,
		Params:   n.Params,
		CorePath: CorePath(n.Core()),
	}

	return json.Marshal(n2)
}

// UnmarshalJSON restores a node from the given JSON data.
// Returns a non-nil error in case of failure.
func (n *Node) UnmarshalJSON(data []byte) error {
	corePath, err := jsonparser.GetString(data, "corePath")
	if err != nil {
		return fmt.Errorf("failed to get core path: %v", err)
	}

	// find the core using the path
	core, ok := WorkingRegistry().GetCore(corePath)
	if !ok {
		return fmt.Errorf("failed to find core path %s in working registry", corePath)
	}

	inputs := Map{}
	eachInput := func(k []byte, v []byte, dType jsonparser.ValueType, offset int) error {
		return restoreDependency(k, v, inputs)
	}

	if err = jsonparser.ObjectEach(data, eachInput, "inputs"); err != nil {
		return fmt.Errorf("failed to unmarshal inputs: %v", err)
	}

	controls := Map{}
	eachControl := func(k []byte, v []byte, dType jsonparser.ValueType, offset int) error {
		return restoreDependency(k, v, controls)
	}

	if err = jsonparser.ObjectEach(data, eachControl, "controls"); err != nil {
		return fmt.Errorf("failed to unmarshal controls: %v", err)
	}

	params := param.Map{}
	eachParam := func(k []byte, v []byte, dType jsonparser.ValueType, offset int) error {
		var p param.Param

		if err := json.Unmarshal(v, &p); err != nil {
			return fmt.Errorf("failed to unmarshal param %s: %v", string(k), err)
		}

		params[string(k)] = &p

		return nil
	}

	if err = jsonparser.ObjectEach(data, eachParam, "params"); err != nil {
		return fmt.Errorf("failed to unmarshal params: %v", err)
	}

	n.core = core
	n.corePath = corePath
	n.Inputs = inputs
	n.Controls = controls
	n.Params = params

	return n.Validate()
}

func (n *Node) Validate() error {
	ifc := n.core.Interface()

	return n.validate(ifc)
}

func (n *Node) validate(ifc *Interface) error {
	if err := ifc.CheckInputs(n.Inputs); err != nil {
		return err
	}

	if err := ifc.CheckParams(n.Params); err != nil {
		return err
	}

	return nil
}

func restoreDependency(key []byte, value []byte, m Map) error {
	var dep Node
	err := json.Unmarshal(value, &dep)
	if err != nil {
		return fmt.Errorf("failed to unmarshal dependency %s: %v", string(key), err)
	}

	m[string(key)] = &dep

	return nil
}
