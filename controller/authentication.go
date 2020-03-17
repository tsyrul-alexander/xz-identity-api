package controller

import (
	"identity-web-api/storage"
	"net/http"
)

type AuthenticationController struct {
	Storage storage.Storage
}

func (controller *AuthenticationController)Hello(w http.ResponseWriter, r *http.Request)  {
	_ = setResponse(w, "Hello")
}