package controller

import (
	"github.com/tsyrul-alexander/identity-web-api/core/authentication"
	"github.com/tsyrul-alexander/identity-web-api/model"
	"github.com/tsyrul-alexander/identity-web-api/model/request"
	"github.com/tsyrul-alexander/identity-web-api/model/response"
	"github.com/tsyrul-alexander/identity-web-api/storage"
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
		setError(w, InvalidRequest, err)
		return
	}
	var user = userRegistration.GetUser()
	if err:= createUser(controller.Storage, user); err != nil {
		setError(w, DbError, err)
		return
	}
	var token, err = controller.Authentication.GenerateToken(user)
	if err != nil {
		setError(w, GenerateTokenError, err)
	}
	SetResponse(w, &response.Registration{Token:token})
}

//

func createUser(dataStorage storage.Storage, user *model.User) error {
	return dataStorage.CreateUser(user)
}