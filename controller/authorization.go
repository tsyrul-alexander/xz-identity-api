package controller

import (
	"identity-web-api/core/authentication"
	"identity-web-api/model"
	"identity-web-api/model/request"
	"identity-web-api/model/response"
	"identity-web-api/storage"
	"net/http"
)

//AuthorizationController ...
type AuthorizationController struct {
	Storage storage.Storage
	Authentication authentication.Authentication
}

//Registration ...
func (controller *AuthorizationController) Registration(w http.ResponseWriter, r *http.Request) {
	var userRegistration = request.UserRegistration{}
	if err := decodeJsonBody(r, &userRegistration); err != nil {
		setError(w, InvalidRequest)
		return
	}
	var user = userRegistration.GetUser()
	if err:= createUser(controller.Storage, user); err != nil {
		setError(w, DbError)
		return
	}
	var token, err = controller.Authentication.GenerateToken(user)
	if err != nil {
		setError(w, GenerateTokenError)
	}
	SetResponse(w, &response.Registration{Token:token})
}

//

func createUser(dataStorage storage.Storage, user *model.User) error {
	return dataStorage.CreateUser(user)
}