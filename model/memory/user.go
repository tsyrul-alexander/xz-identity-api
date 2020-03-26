package memory

import (
	"github.com/google/uuid"
	"github.com/tsyrul-alexander/identity-web-api/model"
)

type User struct {
	Id uuid.UUID
	Roles []model.UserRole
	Tokens []string
}

func CreateUser(id uuid.UUID, roles []model.UserRole) *User {
	return &User{Id:id, Roles:roles}
}
