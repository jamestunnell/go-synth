package node

import (
	"encoding/json"
	"fmt"

	"github.com/buger/jsonparser"
)

// Map is a type alias
type Map = map[string]*Node

// Node provides the framework for running a Core.
type Node struct {
	// Core is the functional core
	Core Core `json:"-"`
	// CorePath is the full core path
	CorePath string `json:"corePath"`
	// Inputs provide functional input to the core every sample.
	Inputs Map `json:"inputs"`
	// Controls provide configuration input to the core once every
	// chunksize samples.
	Controls Map `json:"controls"`
	// Params provide configuration input to the core once at init time
	Params      ParamMap `json:"params"`
	output      *Buffer
	initialized bool
}

// Output returns the latest results from running the core.
// Returns nil before the node is initialized.
func (n *Node) Output() *Buffer {
	return n.output
}

// Initialize initializes the node and its dependencies.
// Dependencies are initialized first.
func (n *Node) Initialize(srate float64, depth int) error {
	ifc := n.Core.Interface()

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

	if err := n.Core.Initialize(args); err != nil {
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

	n.Core.Configure()

	n.Core.Run(n.output)
}

// MarshalJSON generates node JSON data.
// Returns a non-nil error in case of failure.
func (n *Node) MarshalJSON() ([]byte, error) {
	// This needs to be set before we marshal JSON
	n.CorePath = CorePath(n.Core)

	// We marshal an alias that will have all the same fields but none
	// of the methods, so we can avoid an infinite recursion.
	type Alias Node
	return json.Marshal(&struct{ *Alias }{(*Alias)(n)})
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

	params := ParamMap{}
	eachParam := func(k []byte, v []byte, dType jsonparser.ValueType, offset int) error {
		var param interface{}

		if err := json.Unmarshal(v, &param); err != nil {
			return fmt.Errorf("failed to unmarshal param %s: %v", string(k), err)
		}

		params[string(k)] = param

		return nil
	}

	if err = jsonparser.ObjectEach(data, eachParam, "params"); err != nil {
		return fmt.Errorf("failed to unmarshal params: %v", err)
	}

	ifc := core.Interface()

	if err := ifc.CheckInputs(inputs); err != nil {
		return err
	}

	if err := ifc.CheckParams(params); err != nil {
		return err
	}

	n.Core = core
	n.CorePath = corePath
	n.Inputs = inputs
	n.Controls = controls
	n.Params = params

	return nil
}

func (n *Node) Validate() error {
	ifc := n.Core.Interface()

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
