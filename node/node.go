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
	Core Core `json:"core"`
	// CorePath is the full core path
	CorePath string `json:"corePath"`
	// Inputs provide functional input to the core
	Inputs map[string]*Node `json:"inputs"`
	// Controls are used to configure the core
	Controls    map[string]*Node `json:"controls"`
	output      *Buffer
	initialized bool
}

// NewNode makes a new Node.
func NewNode(core Core, inputs, controls map[string]*Node) *Node {
	ifc := core.Interface()
	ifc.EnsureInputs(inputs)
	ifc.EnsureControls(controls)

	return &Node{
		Core:     core,
		Inputs:   inputs,
		Controls: controls,
		CorePath: CorePath(core),
	}
}

// Output returns the latest results from running the core.
// Returns nil before the node is initialized.
func (n *Node) Output() *Buffer {
	return n.output
}

// Initialize initializes the node and its dependencies.
// Dependencies are initialized first.
func (n *Node) Initialize(srate float64, depth int) {
	for _, inputNode := range n.Inputs {
		inputNode.Initialize(srate, depth)
	}

	controlSampleRate := srate / float64(depth)
	for _, controlNode := range n.Controls {
		controlNode.Initialize(controlSampleRate, 1)
	}

	n.Core.Initialize(srate, n.Inputs, n.Controls)

	n.output = NewBuffer(depth)
	n.initialized = true
}

// Run runs the node and its dependencies.
// Dependencies are ran first.
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

// UnmarshalJSON restores a node from the given JSON data.
// Returns a non-nil node in case of failure.
func (n *Node) UnmarshalJSON(data []byte) error {
	corePath, err := jsonparser.GetString(data, "corePath")
	if err != nil {
		return fmt.Errorf("failed to get core path: %v", err)
	}

	// find the core using the path
	core, ok := WorkingRegistry().MakeCore(corePath)
	if !ok {
		return fmt.Errorf("failed to find core path %s in working registry", corePath)
	}

	coreData, _, _, err := jsonparser.Get(data, "core")
	if err != nil {
		return fmt.Errorf("failed to find core: %v", err)
	}

	if err = json.Unmarshal(coreData, &core); err != nil {
		return fmt.Errorf("failed to unmarshal core: %v", err)
	}

	inputs := map[string]*Node{}
	eachInput := func(k []byte, v []byte, dType jsonparser.ValueType, offset int) error {
		return restoreDependency(k, v, inputs)
	}

	if err = jsonparser.ObjectEach(data, eachInput, "inputs"); err != nil {
		return fmt.Errorf("failed to unmarshal inputs: %v", err)
	}

	controls := map[string]*Node{}
	eachControl := func(k []byte, v []byte, dType jsonparser.ValueType, offset int) error {
		return restoreDependency(k, v, controls)
	}

	if err = jsonparser.ObjectEach(data, eachControl, "controls"); err != nil {
		return fmt.Errorf("failed to unmarshal controls: %v", err)
	}

	ifc := core.Interface()
	ifc.EnsureInputs(inputs)
	ifc.EnsureControls(controls)

	n.Core = core
	n.CorePath = corePath
	n.Inputs = inputs
	n.Controls = controls

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
