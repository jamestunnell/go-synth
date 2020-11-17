package typeregistry

import (
	"path"
	"reflect"
)

// TypeRegistry stores known types by their package path.
type TypeRegistry struct {
	pathTypeMap map[string]reflect.Type
}

// New makes an empty type registry
func New() *TypeRegistry {
	return &TypeRegistry{
		pathTypeMap: map[string]reflect.Type{},
	}
}

// TypePath returns the full path of the given type, including package path and type name.
func TypePath(t reflect.Type) string {
	return path.Join(t.PkgPath(), t.Name())
}

// Register adds the given type to the registry.
// Returns false if a type is already registered at the given path.
func (r *TypeRegistry) Register(t reflect.Type) bool {
	path := TypePath(t)
	if _, found := r.pathTypeMap[path]; found {
		return false
	}

	r.pathTypeMap[path] = t

	return true
}

// Unregister removes the given core from the registry.
// Returns false if a type is not registered at the given path.
func (r *TypeRegistry) Unregister(path string) bool {
	if _, found := r.pathTypeMap[path]; !found {
		return false
	}

	delete(r.pathTypeMap, path)

	return true
}

// ForEach calls the given function for each path-type pair in the registry.
func (r *TypeRegistry) ForEach(f func(p string, t reflect.Type)) {
	for path, typ := range r.pathTypeMap {
		f(path, typ)
	}
}

// Paths returns all of the paths for registered types.
func (r *TypeRegistry) Paths() []string {
	paths := make([]string, 0, len(r.pathTypeMap))

	for path := range r.pathTypeMap {
		paths = append(paths, path)
	}

	return paths
}

// GetType uses the given path to look up the registered type.
// Returns false if a type is not found.
func (r *TypeRegistry) GetType(path string) (reflect.Type, bool) {
	t, found := r.pathTypeMap[path]

	return t, found
}

// Clear clears all entries from the registry.
func (r *TypeRegistry) Clear() {
	r.pathTypeMap = map[string]reflect.Type{}
}
