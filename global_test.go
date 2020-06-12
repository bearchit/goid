package goid_test

import (
	"encoding/base64"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/bearchit/goid"
)

func TestGlobalGenerator(t *testing.T) {
	mockGenerator := goid.NewMock(goid.NewUuidV4Generator(true).Generate())
	g := goid.NewGlobalGenerator(
		goid.WithGenerator(mockGenerator),
	)

	id := g.Prefix("object").Generate()

	decoded, err := base64.URLEncoding.DecodeString(id.String())
	require.NoError(t, err)

	assert.Equal(t, fmt.Sprintf("object:%s", mockGenerator.Generate()), string(decoded))
}

func TestParseGlobalID(t *testing.T) {
	mockGenerator := goid.NewMock(goid.NewUuidV4Generator(true).Generate())
	g := goid.NewGlobalGenerator(
		goid.WithGenerator(mockGenerator),
	)

	id := g.Prefix("object").Generate()
	mockedID := mockGenerator.Generate()

	globalID, err := goid.ParseGlobalID(id)
	require.NoError(t, err)

	assert.Equal(t, "object", globalID.Prefix)
	assert.Equal(t, ":", globalID.Delimiter)
	assert.Equal(t, mockedID, globalID.ID)
}
