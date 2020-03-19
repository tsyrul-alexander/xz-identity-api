package request

import (
	"github.com/google/uuid"
	"identity-web-api/model"
)

//UserRegistration ...
type UserRegistration struct {
	Login string `json:"login"`
	Password string `json:"password"`
	Name string `json:"name"`
}

func (userRegistration *UserRegistration) GetUser() *model.User {
	return &model.User{
		ID:           uuid.New(),
		Name:         userRegistration.Name,
		IdentityType: model.IdentityTypeDefault,
		DefaultIdentity: model.DefaultIdentity{
			ID:       uuid.New(),
			Login:    userRegistration.Login,
			Password: model.CreateHashPassword(userRegistration.Password),
		},
	}
}
