package goid_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/bearchit/goid"
)

func TestNamespaceGenerator_Generate(t *testing.T) {
	g := goid.NewNamespaceGenerator("com")
	id := g.Generate("github", "bearchit", "goid")
	assert.Equal(t, "com.github.bearchit.goid", id.String())
}

func TestWithDelimiter(t *testing.T) {
	g := goid.NewNamespaceGenerator("com", goid.WithDelimiter("-"))
	id := g.Generate("github", "bearchit", "goid")
	assert.Equal(t, "com-github-bearchit-goid", id.String())
}
