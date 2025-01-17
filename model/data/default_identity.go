package data

import (
	"github.com/google/uuid"
	"github.com/tsyrul-alexander/xz-identity-api/model"
)

type DefaultIdentity struct {
	ID       uuid.UUID
	Login    string
	Password model.HashPassword
}
