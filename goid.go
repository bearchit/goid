package goid

type ID string

func (id ID) String() string {
	return string(id)
}

func FromString(id string) ID {
	return ID(id)
}

type IDs []ID

func NewIDs(id ...ID) IDs {
	ids := make(IDs, 0)
	ids = append(ids, id...)
	return ids
}

func (ids IDs) Contains(id ID) bool {
	for _, x := range ids {
		if x == id {
			return true
		}
	}
	return false
}

type Generator interface {
	Generate() ID
}
