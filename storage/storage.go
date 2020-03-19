package storage

import (
	"identity-web-api/model"
)

//Storage ...
type Storage interface {
	CreateUser(user *model.User) error
	GetUser(login string) (*model.User, error)
}
