package storage

import (
	"identity-web-api/model"
)

//Storage ...
type Storage interface {
	CreateUser(user *model.User) error
}
