package storage

import (
	"github.com/tsyrul-alexander/identity-web-api/model"
)

//Storage ...
type Storage interface {
	CreateUser(user *model.User) error
	GetUser(login string) (*model.User, error)
}
