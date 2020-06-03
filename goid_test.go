package goid_test

import (
	"testing"

	"github.com/bearchit/goid"
	"github.com/stretchr/testify/assert"
)

func TestID_Nil(t *testing.T) {
	t.Run("with blank string", func(t *testing.T) {
		id := goid.ID("")
		assert.True(t, id.Nil())
	})

	t.Run("with NilID constant", func(t *testing.T) {
		id := goid.NilID
		assert.True(t, id.Nil())
	})
}

func TestFromString(t *testing.T) {
	id := goid.FromString("hello")
	assert.Equal(t, goid.ID("hello"), id)
}

func TestNewIDs(t *testing.T) {
	ids := goid.NewIDs("1", "2", "3", "4")
	assert.Equal(t, goid.ID("1"), ids[0])
	assert.Equal(t, 4, len(ids))
}

func TestIDs_Contains(t *testing.T) {
	ids := goid.NewIDs("1", "2", "3", "4")
	assert.True(t, ids.Contains("1"))
	assert.False(t, ids.Contains("5"))
}
