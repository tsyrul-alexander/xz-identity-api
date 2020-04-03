package controller

import (
	"github.com/tsyrul-alexander/xz-identity-api/core/authentication"
	"github.com/tsyrul-alexander/xz-identity-api/model/memory"
	"github.com/tsyrul-alexander/xz-identity-api/model/response"
	"github.com/tsyrul-alexander/xz-identity-api/storage"
	"net/http"
)

type AuthenticationController struct {
	DataStorage    storage.DataStorage
	MemoryStorage  storage.MemoryStorage
	Authentication authentication.Authentication
}

func (controller *AuthenticationController)Login(w http.ResponseWriter, r *http.Request) {
	var query = r.URL.Query()
	var login = query.Get("login")
	var password = query.Get("password")
	if login == "" || password == "" {
		setError(w, InvalidRequest, nil)
		return
	}
	var user, err = controller.DataStorage.GetUserByLogin(login)
	if err != nil {
		setError(w, DbError, err)
		return
	}
	if user == nil || !user.DefaultIdentity.Password.GetIsCompareHashPassword(password) {
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

func (controller *AuthenticationController)GetUserInfo(w http.ResponseWriter, r *http.Request) {
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