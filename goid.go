package goid

type ID string

const NilID ID = ""

func (id ID) Nil() bool {
	return id == NilID
}

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

func (ids *IDs) Add(others ...ID) {
	*ids = append(*ids, others...)
}

func (ids IDs) Contains(others ...ID) bool {
	for _, x := range ids {
		for _, other := range others {
			if x == other {
				return true
			}
		}
	}
	return false
}

type Generator interface {
	Generate() ID
}
