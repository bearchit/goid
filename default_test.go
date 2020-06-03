package goid_test

import (
	"github.com/bearchit/goid"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_DefaultGenerator_Generate(t *testing.T) {
	_, err := uuid.FromString(goid.Generate().String())
	require.NoError(t, err)
}

func TestDefaultMockGenerator_Generate(t *testing.T) {
	g := goid.DefaultMockGenerator
	id := g.Generate()
	assert.Equal(t, goid.ID("mocked-id"), id)
}
