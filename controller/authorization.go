package controller

import (
	"github.com/google/uuid"
	"identity-web-api/model"
	"identity-web-api/storage"
	"net/http"
)

//AuthorizationController ...
type AuthorizationController struct {
	Storage storage.Storage
}

//Registration ...
func (controller *AuthorizationController) Registration(w http.ResponseWriter, r *http.Request) {
	var userRegistration = model.UserRegistration{}
	var err = decodeJsonBody(r, &userRegistration)
	if err == nil {
		err = createUser(controller.Storage, &userRegistration)
	}
	if err != nil {
		err = setResponse(w, err.Error())
	}
	err = setResponse(w, "Ok")
	err = r.Body.Close()
}
func createUser(dataStorage storage.Storage, userRegistration *model.UserRegistration) error {
	var user = model.User{
		ID:           uuid.New(),
		Name:         userRegistration.Name,
		IdentityType: model.IdentityTypeDefault,
		DefaultIdentity: model.DefaultIdentity{
			Login:    userRegistration.Login,
			Password: userRegistration.Password,
		},
	}
	return dataStorage.CreateUser(&user)
}