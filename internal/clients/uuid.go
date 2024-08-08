package clients

import "github.com/google/uuid"

type IDGenerator interface {
	GetID() string
}

type uuidGenerator struct{}

func NewUuidGenerator() IDGenerator {
	return &uuidGenerator{}
}

func (u *uuidGenerator) GetID() string {
	return uuid.NewString()
}
