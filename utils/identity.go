package utils

import "github.com/google/uuid"

type IdentityUtil interface {
	NewId() string
}

type UuidIdentity struct {
}

func (*UuidIdentity) NewId() string {
	return uuid.NewString()
}

func NewUuidIdentity() *UuidIdentity {
	return &UuidIdentity{}
}
