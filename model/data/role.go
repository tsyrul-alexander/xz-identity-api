package data

import (
	"github.com/google/uuid"
	"github.com/tsyrul-alexander/identity-web-api/model"
)

type Role struct {
	Id uuid.UUID
	Name string
	Code model.UserRole
}
