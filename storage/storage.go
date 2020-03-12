package storage

import (
	"../model"
)

//Storage ...
type Storage interface {
	createUser(user *model.User) error
}
