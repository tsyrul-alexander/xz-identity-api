package controller

import (
	"identity-web-api/core/authentication"
	"identity-web-api/model/request"
	"identity-web-api/model/response"
	"identity-web-api/storage"
	"net/http"
)

type AuthenticationController struct {
	Storage storage.Storage
	Authentication authentication.Authentication
}

func (controller *AuthenticationController)Login(w http.ResponseWriter, r *http.Request)  {
	var userLogin = &request.UserLogin{}
	if err := decodeJsonBody(r, &userLogin); err != nil {
		setError(w, InvalidRequest)
		return
	}
	var user, err = controller.Storage.GetUser(userLogin.Login)
	if err != nil {
		setError(w, DbError)
		return
	}
	if !user.DefaultIdentity.Password.GetIsCompareHashPassword(userLogin.Password) {
		setError(w, InvalidCredential)
		return
	}
	var token, e = controller.Authentication.GenerateToken(user)
	if e != nil {
		setError(w, GenerateTokenError)
	}
	SetResponse(w, response.Login{Token:token})
}