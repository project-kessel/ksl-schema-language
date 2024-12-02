package ksl

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBad(t *testing.T) {
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
