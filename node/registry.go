package node

import (
	"reflect"

	"github.com/jamestunnell/go-synth/util/typeregistry"
)

// CoreRegistry maps core types to their paths
type CoreRegistry struct {
	typeReg *typeregistry.TypeRegistry
}

var registry = NewCoreRegistry()

// NewCoreRegistry makes a new instance.
func NewCoreRegistry() *CoreRegistry {
	return &CoreRegistry{typeReg: typeregistry.New()}
}

// WorkingRegistry returns the working registry
func WorkingRegistry() *CoreRegistry {
	return registry
}

// CorePath returns the full core path, including package path and core type.
func CorePath(c Core) string {
	return typeregistry.TypePath(reflect.TypeOf(c).Elem())
}

// Register adds the given core to the registry.
func (r *CoreRegistry) Register(c Core) bool {
	return r.typeReg.Register(reflect.TypeOf(c).Elem())
}

// Unregister removes the type at the given path from the registry.
func (r *CoreRegistry) Unregister(path string) bool {
	return r.typeReg.Unregister(path)
}

// Paths returns all of the paths for registered cores.
func (r *CoreRegistry) Paths() []string {
	return r.typeReg.Paths()
}

// GetCore uses the given path to look up the registered core type
// and then produce a new core instance.
// Returns false if a type is not found.
func (r *CoreRegistry) GetCore(path string) (Core, bool) {
	t, found := r.typeReg.GetType(path)
	if !found {
		return nil, false
	}

	c, ok := reflect.New(t).Interface().(Core)

	return c, ok
}
