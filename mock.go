package goid

type MockGenerator struct {
	id ID
}

func NewMock(id ID) Generator {
	return &MockGenerator{id: id}
}

func (g MockGenerator) Generate(_ ...string) ID {
	return g.id
}
