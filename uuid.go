package goid

import uuid "github.com/satori/go.uuid"

type uuidV4Generator struct {
}

func NewUuidV4Generator() Generator {
	return &uuidV4Generator{}
}

func (g uuidV4Generator) Generate() ID {
	return ID(uuid.NewV4().String())
}
