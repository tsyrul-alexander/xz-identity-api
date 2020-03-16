package controller

import (
	"../model"
	"../server"
	"../storage"
	"encoding/json"
	"github.com/google/uuid"
	"net/http"
)

type AuthorizationController struct {
	Server *server.Server
}

func (controller *AuthorizationController) Registration(w http.ResponseWriter, r *http.Request) {
	var userRegistration = model.UserRegistration{}
	var err = decodeJsonBody(r, &userRegistration)
	if err == nil {
		err = createUser(controller.Server.Storage, &userRegistration)
	}
	if err != nil {
		err = setResponse(w, err.Error())
	}
	err = setResponse(w, "Ok")
	err = r.Body.Close()
}
func createUser(dataStorage storage.Storage, userRegistration *model.UserRegistration) error {
	var user = model.User{
		ID: uuid.New(),
		Name: userRegistration.Name,
		IdentityType: model.IdentityTypeDefault,
		DefaultIdentity: model.DefaultIdentity{
			Login:    userRegistration.Login,
			Password: userRegistration.Password,
		},
	}
	return dataStorage.CreateUser(&user)
}
func setResponse(w http.ResponseWriter, data interface{}) error {
	w.Header().Add("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(data)
}

func decodeJsonBody(r *http.Request, obj interface{}) error {
	var jsonDecoder = json.NewDecoder(r.Body)
	return jsonDecoder.Decode(obj)
}