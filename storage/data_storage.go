package storage

import (
	"github.com/google/uuid"
	"github.com/tsyrul-alexander/identity-web-api/model"
	"github.com/tsyrul-alexander/identity-web-api/model/data"
)

//DataStorage ...
type DataStorage interface {
	CreateUser(user *data.User, roles ...model.UserRole) error
	CreateUserRole(userId uuid.UUID, roles ...model.UserRole) error
	GetUserByLogin(login string) (*data.User, error)
	GetUserById(id uuid.UUID) (*data.User, error)
	GetUserRoles(id uuid.UUID) ([]model.UserRole, error)
}
