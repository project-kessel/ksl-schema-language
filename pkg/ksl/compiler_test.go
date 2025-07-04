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
