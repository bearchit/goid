package goid

import uuid "github.com/satori/go.uuid"

type ID string

func (id ID) String() string {
	return string(id)
}

type Generator interface {
	Generate() ID
}

type uuidV4Generator struct {
}

func NewUuidV4Generator() Generator {
	return &uuidV4Generator{}
}

func (g uuidV4Generator) Generate() ID {
	return ID(uuid.NewV4().String())
}

type MockGenerator struct {
	id ID
}

func NewMock(id ID) Generator {
	return &MockGenerator{id: id}
}

func (g MockGenerator) Generate() ID {
	return g.id
}

var (
	DefaultGenerator     = NewUuidV4Generator()
	MockedID             = ID("mocked-id")
	DefaultMockGenerator = MockGenerator{id: MockedID}
)

func Generate() ID {
	return DefaultGenerator.Generate()
}
