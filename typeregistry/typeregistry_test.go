package typeregistry_test

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jamestunnell/go-synth/typeregistry"
)

const trPath = "github.com/jamestunnell/go-synth/typeregistry/TypeRegistry"

func TestNew(t *testing.T) {
	r := typeregistry.New()

	assert.NotNil(t, r)
}

func TestTypePath(t *testing.T) {
	r := typeregistry.New()
	path := typeregistry.TypePath(reflect.TypeOf(r).Elem())

	assert.Equal(t, path, trPath)
}

func TestTypeRegistryBasicUsage(t *testing.T) {
	r := typeregistry.New()
	typ := reflect.TypeOf(r).Elem()
	path := typeregistry.TypePath(typ)

	assert.True(t, r.Register(typ))
	assert.False(t, r.Register(typ))

	typ2, found := r.GetType(path)

	assert.True(t, found)
	assert.Equal(t, typ, typ2)

	assert.True(t, r.Unregister(path))
	assert.False(t, r.Unregister(path))

	_, found = r.GetType(path)

	assert.False(t, found)
}

func TestTypeRegistryMultipleTypes(t *testing.T) {
	r := typeregistry.New()
	typ1 := reflect.TypeOf(r).Elem()
	typ2 := reflect.TypeOf(5)
	typ3 := reflect.TypeOf("abc")
	expectedTypes := []reflect.Type{typ1, typ2, typ3}
	path1 := typeregistry.TypePath(typ1)
	path2 := typeregistry.TypePath(typ2)
	path3 := typeregistry.TypePath(typ3)
	expectedPaths := []string{path1, path2, path3}

	assert.True(t, r.Register(typ1))
	assert.True(t, r.Register(typ2))
	assert.True(t, r.Register(typ3))

	paths := r.Paths()

	assert.Equal(t, expectedPaths, paths)

	recvdPaths := []string{}
	recvdTypes := []reflect.Type{}

	r.ForEach(func(path string, typ reflect.Type) {
		recvdPaths = append(recvdPaths, path)
		recvdTypes = append(recvdTypes, typ)
	})

	assert.Equal(t, expectedPaths, recvdPaths)
	assert.Equal(t, expectedTypes, recvdTypes)

	r.Clear()

	paths = r.Paths()
	assert.Empty(t, paths)
}
