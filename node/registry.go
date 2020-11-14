package node

import (
	"log"
	"path"
	"reflect"
)

// Registry stores known core types by their package path.
type Registry struct {
	pathTypeMap map[string]reflect.Type
}

var registry = NewRegistry()

// NewRegistry makes an empty registry
func NewRegistry() *Registry {
	return &Registry{
		pathTypeMap: map[string]reflect.Type{},
	}
}

// WorkingRegistry returns the working registry
func WorkingRegistry() *Registry {
	return registry
}

// CorePath returns the full core path, including package path and core type.
func CorePath(c Core) string {
	return TypePath(reflect.TypeOf(c).Elem())
}

// TypePath returns the full path of the given type, including package path and type name.
func TypePath(t reflect.Type) string {
	return path.Join(t.PkgPath(), t.Name())
}

// RegisterCore adds the given core to the registry.
func (r *Registry) RegisterCore(c Core) {
	t := reflect.TypeOf(c).Elem()
	path := TypePath(t)

	if _, found := r.pathTypeMap[path]; !found {
		log.Printf("registering core %s", path)

		r.pathTypeMap[path] = t
	}
}

// UnregisterCore removes the given core from the registry.
func (r *Registry) UnregisterCore(c Core) {
	t := reflect.TypeOf(c).Elem()
	path := TypePath(t)

	if _, found := r.pathTypeMap[path]; found {
		log.Printf("unregistering core %s", path)

		delete(r.pathTypeMap, path)
	}
}

// Paths returns all of the paths for registered cores.
func (r *Registry) Paths() []string {
	paths := make([]string, 0, len(r.pathTypeMap))

	for path := range r.pathTypeMap {
		paths = append(paths, path)
	}

	return paths
}

// MakeCore uses the given path to look among the registered core types.
// If found, the type is used to make a new instance.
// Returns false if the type is not found.
func (r *Registry) MakeCore(path string) (Core, bool) {
	t, found := r.pathTypeMap[path]
	if !found {
		return nil, false
	}

	return reflect.New(t).Interface().(Core), true
}
