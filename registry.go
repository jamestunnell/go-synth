package synth

type MakeBlockFunc func() Block

// BlockRegistry maps core types to their paths
type BlockRegistry struct {
	makeFuncs map[string]MakeBlockFunc
}

var registry = NewBlockRegistry()

// NewBlockRegistry makes a new instance.
func NewBlockRegistry() *BlockRegistry {
	return &BlockRegistry{makeFuncs: map[string]MakeBlockFunc{}}
}

// WorkingRegistry returns the working registry
func WorkingRegistry() *BlockRegistry {
	return registry
}

// Register adds the given core to the registry.
func (r *BlockRegistry) Register(f MakeBlockFunc) bool {
	b := f()
	path := BlockPath(b)

	if _, found := r.makeFuncs[path]; found {
		return false
	}

	r.makeFuncs[path] = f

	return true
}

// Unregister removes the type at the given path from the registry.
func (r *BlockRegistry) Unregister(path string) bool {
	if _, found := r.makeFuncs[path]; !found {
		return false
	}

	delete(r.makeFuncs, path)

	return true
}

// Paths returns all of the paths for registered cores.
func (r *BlockRegistry) Paths() []string {
	paths := []string{}

	for path := range r.makeFuncs {
		paths = append(paths, path)
	}

	return paths
}

// MakeBlock uses the given path to look up a make function and
// create a new block.
// Returns false if no make function is registered.
func (r *BlockRegistry) MakeBlock(path string) (Block, bool) {
	f, found := r.makeFuncs[path]
	if !found {
		return nil, false
	}

	return f(), true
}
