package controller

import (
	"github.com/tsyrul-alexander/identity-web-api/core/authentication"
	"github.com/tsyrul-alexander/identity-web-api/model/request"
	"github.com/tsyrul-alexander/identity-web-api/model/response"
	"github.com/tsyrul-alexander/identity-web-api/storage"
	"net/http"
)

type AuthenticationController struct {
	Storage storage.Storage
	Authentication authentication.Authentication
}

func (controller *AuthenticationController)Login(w http.ResponseWriter, r *http.Request)  {
	var userLogin = &request.UserLogin{}
	if err := decodeJsonBody(r, &userLogin); err != nil {
		setError(w, InvalidRequest, err)
		return
	}
	var user, err = controller.Storage.GetUser(userLogin.Login)
	if err != nil {
		setError(w, DbError, err)
		return
	}
	if user == nil || !user.DefaultIdentity.Password.GetIsCompareHashPassword(userLogin.Password) {
		setError(w, InvalidCredential, nil)
		return
	}
	var token, e = controller.Authentication.GenerateToken(user)
	if e != nil {
		setError(w, GenerateTokenError, err)
	}
	SetResponse(w, response.Login{Token:token})
}