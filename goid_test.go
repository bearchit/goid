package goid_test

import (
	"testing"

	"github.com/bearchit/goid"
	"github.com/stretchr/testify/assert"
)

func TestFromString(t *testing.T) {
	id := goid.FromString("hello")
	assert.Equal(t, goid.ID("hello"), id)
}
