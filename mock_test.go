package goid_test

import (
	"github.com/bearchit/goid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_MockGenerator_Generate(t *testing.T) {
	g := goid.NewMock("fixed-id-for-mock")
	assert.Equal(t, "fixed-id-for-mock", g.Generate().String())
}
