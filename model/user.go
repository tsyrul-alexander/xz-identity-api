package model

import "github.com/google/uuid"

//User ...
type User struct {
	ID uuid.UUID
	Name string
	IdentityType IdentityType
	DefaultIdentity DefaultIdentity
}

//Create ...
func Create() *User {
	var id = uuid.New()
	return &User{ID: id}
}
