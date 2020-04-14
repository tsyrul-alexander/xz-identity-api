package controller

import (
	"github.com/tsyrul-alexander/xz-identity-api/core/authentication"
	"github.com/tsyrul-alexander/xz-identity-api/model"
	"github.com/tsyrul-alexander/xz-identity-api/model/memory"
	"github.com/tsyrul-alexander/xz-identity-api/model/response"
	"github.com/tsyrul-alexander/xz-identity-api/storage"
	"log"
	"net/http"
	"strconv"
)

type AuthenticationController struct {
	DataStorage    storage.DataStorage
	MemoryStorage  storage.MemoryStorage
	Authentication authentication.Authentication
}

func (controller *AuthenticationController) Login(w http.ResponseWriter, r *http.Request) {
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

func (controller *AuthenticationController) GetUserInRoles(w http.ResponseWriter, r *http.Request) {
	var query = r.URL.Query()
	var token = query.Get("token")
	var rolesStr = query["role"]
	if len(rolesStr) == 0 || token == "" {
		setError(w, InvalidRequest, nil)
	}
	var userInfo, err = controller.getUserInfo(token)
	if err != nil {
		setError(w, err, nil)
	}
	for _, roleStr := range rolesStr {
		var roleInt, parseErr = strconv.Atoi(roleStr)
		if parseErr != nil {
			setError(w, InvalidRequest, parseErr)
		}
		if userInfo.GetIfExistRole(model.UserRole(roleInt)) {
			SetResponse(w, &response.IdentityResponse {Success:true})
			return
		}
	}
	SetResponse(w, &response.IdentityResponse{Success:false})
}

func (controller *AuthenticationController) GetUserInfo(w http.ResponseWriter, r *http.Request) {
	var token = getAuthorizedToken(r)
	if token == "" {
		setError(w, AuthenticationRequired, nil)
	}
	var userInfo, err = controller.getUserInfo(token)
	if err != nil {
		setError(w, err, nil)
	}
	SetResponse(w, userInfo)
}

func (controller *AuthenticationController) getUserInfo(token string) (*response.UserInfo, *ServiceError) {
	var userId, tokenErr = controller.Authentication.GetUserId(token)
	if tokenErr != nil {
		log.Println(tokenErr.Error())
		return nil, ParseTokenError
	}
	if user, exist := controller.MemoryStorage.GetUser(userId); exist {
		return response.CreateUserInfo(user.Id, user.Roles), nil
	}
	var user, dbErr = controller.DataStorage.GetUserRoles(userId)
	if dbErr != nil {
		log.Println(dbErr.Error())
		return nil, DbError
	}
	return response.CreateUserInfo(userId, user), nil
}