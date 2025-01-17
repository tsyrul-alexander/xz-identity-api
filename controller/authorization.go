package controller

import (
	"github.com/tsyrul-alexander/xz-identity-api/core/authentication"
	"github.com/tsyrul-alexander/xz-identity-api/model"
	"github.com/tsyrul-alexander/xz-identity-api/model/data"
	"github.com/tsyrul-alexander/xz-identity-api/model/request"
	"github.com/tsyrul-alexander/xz-identity-api/model/response"
	"github.com/tsyrul-alexander/xz-identity-api/storage"
	"net/http"
)

//AuthorizationController ...
type AuthorizationController struct {
	Storage storage.DataStorage
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

func createUser(dataStorage storage.DataStorage, user *data.User) error {
	return  dataStorage.CreateUser(user, model.UserRoleClient)
}