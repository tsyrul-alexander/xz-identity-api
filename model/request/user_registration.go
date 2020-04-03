package request

import (
	"github.com/google/uuid"
	"github.com/tsyrul-alexander/xz-identity-api/model"
	"github.com/tsyrul-alexander/xz-identity-api/model/data"
)

//UserRegistration ...
type UserRegistration struct {
	Login string `json:"login"`
	Password string `json:"password"`
	Name string `json:"name"`
}

func (userRegistration *UserRegistration) GetUser() *data.User {
	return &data.User{
		ID:           uuid.New(),
		Name:         userRegistration.Name,
		IdentityType: data.IdentityTypeDefault,
		DefaultIdentity: data.DefaultIdentity{
			ID:       uuid.New(),
			Login:    userRegistration.Login,
			Password: model.CreateHashPassword(userRegistration.Password),
		},
	}
}
