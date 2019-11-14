package goid

import uuid "github.com/satori/go.uuid"

type ID string

func (id ID) String() string {
	return string(id)
}

type Generator interface {
	Generate() ID
}

var DefaultGenerator = defaultGenerator{}

type defaultGenerator struct {
}

func (g defaultGenerator) Generate() ID {
	return ID(uuid.NewV4().String())
}

var (
	MockedID             = ID("1")
	DefaultMockGenerator = MockGenerator{id: MockedID}
)

type MockGenerator struct {
	id ID
}

func NewMock(id ID) Generator {
	return &MockGenerator{id: id}
}

func (g MockGenerator) Generate() ID {
	return g.id
}

func Generate() ID {
	return DefaultGenerator.Generate()
}
