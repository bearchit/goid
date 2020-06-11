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
	mockedID := goid.NewMock(goid.NewUuidV4Generator(true).Generate()).Generate()
	g := goid.NewGlobalGenerator()

	id := g.Prefix("object").ID(mockedID).Generate()

	decoded, err := base64.URLEncoding.DecodeString(id.String())
	require.NoError(t, err)

	assert.Equal(t, fmt.Sprintf("object:%s", mockedID), string(decoded))
}
