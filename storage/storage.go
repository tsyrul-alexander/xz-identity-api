package storage

import (
	"../model"
)

//Storage ...
type Storage interface {
	CreateUser(user *model.User) error
}
