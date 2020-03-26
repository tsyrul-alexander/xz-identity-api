package controller

import (
	"github.com/tsyrul-alexander/identity-web-api/core/authentication"
	"github.com/tsyrul-alexander/identity-web-api/model/memory"
	"github.com/tsyrul-alexander/identity-web-api/model/request"
	"github.com/tsyrul-alexander/identity-web-api/model/response"
	"github.com/tsyrul-alexander/identity-web-api/storage"
	"net/http"
)

type AuthenticationController struct {
	DataStorage    storage.DataStorage
	MemoryStorage  storage.MemoryStorage
	Authentication authentication.Authentication
}

func (controller *AuthenticationController)Login(w http.ResponseWriter, r *http.Request) {
	var userLogin = &request.UserLogin{}
	if err := decodeJsonBody(r, &userLogin); err != nil {
		setError(w, InvalidRequest, err)
		return
	}
	var user, err = controller.DataStorage.GetUserByLogin(userLogin.Login)
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
	var userRoles, _ = controller.DataStorage.GetUserRoles(user.ID)
	controller.MemoryStorage.SetUser(memory.CreateUser(user.ID, userRoles))
	SetResponse(w, response.Login{Token:token})
}

func (controller *AuthenticationController)GetUserInfo(w http.ResponseWriter, r *http.Request)  {
	var token = getAuthorizedToken(r)
	if token == "" {
		setError(w, AuthenticationRequired, nil)
	}
	var userId, tokenErr = controller.Authentication.GetUserId(token)
	if tokenErr != nil {
		setError(w, ParseTokenError, tokenErr)
	}
	if user, exist := controller.MemoryStorage.GetUser(userId); exist {
		SetResponse(w, response.CreateUserInfo(user.Id, user.Roles))
		return
	}
	var user, dbErr = controller.DataStorage.GetUserById(userId)
	if dbErr != nil {
		setError(w, DbError, dbErr)
		return
	}
	SetResponse(w, user)
}