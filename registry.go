package synth

import (
	"reflect"

	"github.com/jamestunnell/go-synth/util/typeregistry"
)

// BlockRegistry maps core types to their paths
type BlockRegistry struct {
	typeReg *typeregistry.TypeRegistry
}

var registry = NewBlockRegistry()

// NewBlockRegistry makes a new instance.
func NewBlockRegistry() *BlockRegistry {
	return &BlockRegistry{typeReg: typeregistry.New()}
}

// WorkingRegistry returns the working registry
func WorkingRegistry() *BlockRegistry {
	return registry
}

// BlockPath returns the full core path, including package path and core type.
func BlockPath(c Block) string {
	return typeregistry.TypePath(reflect.TypeOf(c).Elem())
}

// Register adds the given core to the registry.
func (r *BlockRegistry) Register(c Block) bool {
	return r.typeReg.Register(reflect.TypeOf(c).Elem())
}

// Unregister removes the type at the given path from the registry.
func (r *BlockRegistry) Unregister(path string) bool {
	return r.typeReg.Unregister(path)
}

// Paths returns all of the paths for registered cores.
func (r *BlockRegistry) Paths() []string {
	return r.typeReg.Paths()
}

// GetBlock uses the given path to look up the registered core type
// and then produce a new core instance.
// Returns false if a type is not found.
func (r *BlockRegistry) GetBlock(path string) (Block, bool) {
	t, found := r.typeReg.GetType(path)
	if !found {
		return nil, false
	}

	c, ok := reflect.New(t).Interface().(Block)

	return c, ok
}
