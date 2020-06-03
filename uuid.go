package goid

import (
	"encoding/hex"
	uuid "github.com/satori/go.uuid"
)

type uuidV4Generator struct {
	dash bool
}

func NewUuidV4Generator(dash bool) Generator {
	return &uuidV4Generator{
		dash: dash,
	}
}

func (g uuidV4Generator) Generate() ID {
	id := uuid.NewV4()

	if g.dash {
		return ID(id.String())
	}

	return ID(hex.EncodeToString(id.Bytes()))
}
