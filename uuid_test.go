package goid_test

import (
	"github.com/bearchit/goid"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestUuidV4Generator_Generate(t *testing.T) {
	t.Run("when using dash", func(t *testing.T) {
		g := goid.NewUuidV4Generator(true)
		id := g.Generate()
		assert.True(t, strings.Contains(id.String(), "-"))
	})

	t.Run("when not using dash", func(t *testing.T) {
		g := goid.NewUuidV4Generator(false)
		id := g.Generate()
		assert.False(t, strings.Contains(id.String(), "-"))
	})
}
