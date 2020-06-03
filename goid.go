package goid

type ID string

func (id ID) String() string {
	return string(id)
}

func FromString(id string) ID {
	return ID(id)
}

type Generator interface {
	Generate() ID
}
