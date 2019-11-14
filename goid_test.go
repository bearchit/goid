package goid_test

import (
	"testing"

	"github.com/bearchit/goid"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_DefaultGenerator_Generate(t *testing.T) {
	_, err := uuid.FromString(goid.Generate().String())
	require.NoError(t, err)
}

func Test_MockGenerator_Generate(t *testing.T) {
	g := goid.NewMock("fixed-id-for-mock")
	assert.Equal(t, "fixed-id-for-mock", g.Generate().String())
}
