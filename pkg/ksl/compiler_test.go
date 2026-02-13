package ksl

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrorsAreThrownWhenSyntaxIsIncorrect(t *testing.T) {
	_, err := Compile(strings.NewReader(`
version 
namespace rbac

public type  {}

public type group {
    relation member: [Any principal or group.membr]
}`))

	assert.ErrorIs(t, err, ErrSyntaxError)
}

func TestCompilerPassesWithCorrectlyFormattedFile(t *testing.T) {
	ns, err := Compile(strings.NewReader(`
version 0.1
namespace rbac

public type principal {}

public type group {
    relation member: [Any principal or group.membr]
}`))

	assert.NoError(t, err)
	assert.Equal(t, "rbac", ns.Name)
}

func TestCompilerErrorsWithIncorrectStartOfFile(t *testing.T) {
	_, err := Compile(strings.NewReader(`
versions 0.1
namespaced rbac
public type principal {}
public type group {
    relation member: [Any principal or group.membr]
}`))

	assert.Error(t, err)
}

func TestCompilerPassesWithVersionRelationship(t *testing.T) {
	ns, err := Compile(strings.NewReader(`
version 0.1
namespace rbac
public type principal {}
public type group {
    relation member: [Any principal or group.membr]
	relation #version: [Any principal]
}`))

	assert.NoError(t, err)
	assert.Equal(t, "rbac", ns.Name)
}

func TestCompilerFailsWithVersionRelationship(t *testing.T) {
	_, err := Compile(strings.NewReader(`
version 0.1
namespace rbac
public type principal {}
public type group {
    relation member: [Any principal or group.membr]
	relation version: [Any principal]
}`))

	assert.Error(t, err)
}

func TestCompilerPassesWithNamespaceRelationship(t *testing.T) {
	ns, err := Compile(strings.NewReader(`
version 0.1
namespace rbac
public type principal {}
public type group {
    relation member: [Any principal or group.membr]
	relation #namespace: [Any principal]
}`))

	assert.NoError(t, err)
	assert.Equal(t, "rbac", ns.Name)
}

func TestCompilerPassesWithExtendedType(t *testing.T) {
	ns, err := Compile(strings.NewReader(`
version 0.1
namespace rbac
public type principal {}
public type group {
    relation member: [Any principal or group.membr]
	relation #namespace: [Any principal]
}
extend type group {
	relation new_rel: member
}`))

	assert.NoError(t, err)
	assert.Equal(t, 1, len(ns.ExtendedTypes))
	assert.Equal(t, "group", ns.ExtendedTypes[0].Name)
}

func TestCompilerPassesWithExtendedTypeAndWithExtension(t *testing.T) {
	ns, err := Compile(strings.NewReader(`
version 0.1
namespace rbac
public type principal {}
public type group {
    relation member: [Any principal or group.membr]
	relation #namespace: [Any principal]
}
extend type group {
	relation new_rel: member with fish(fin: 'sharp', knee: 'pop')
}`))

	assert.NoError(t, err)
	assert.Equal(t, 1, len(ns.ExtendedTypes))
	assert.Equal(t, "group", ns.ExtendedTypes[0].Name)
	assert.Equal(t, 1, len(ns.ExtensionReferences))
	assert.Equal(t, "fish", ns.ExtensionReferences[0].Name)
	assert.Equal(t, 2, len(ns.ExtensionReferences[0].Params))
	assert.Equal(t, "sharp", ns.ExtensionReferences[0].Params["fin"])
	assert.Equal(t, "pop", ns.ExtensionReferences[0].Params["knee"])
}
