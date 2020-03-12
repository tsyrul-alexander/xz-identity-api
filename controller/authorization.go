package controller

import (
	"../model"
	"../storage"
	"encoding/json"
	"net/http"
)

func Registration(w http.ResponseWriter, r *http.Request) {
	var user = model.User{}
	var err = decodeJsonBody(r, &user)
	if err == nil {
		err = createPQStore(&user)
	}
	if err != nil {
		err = setResponse(w, err.Error())
	}
	err = setResponse(w, "Ok")
	err = r.Body.Close()
}

func setResponse(w http.ResponseWriter, data interface{}) error {
	w.Header().Add("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(data)
}

func createPQStore(user *model.User) error  {
	var config = storage.Config{ConnectionString: "user=postgres password=123 dbname=Test sslmode=disable"}
	var pqStorage = storage.CreatePQStore(&config)
	return pqStorage.CreateUser(user)
}

func decodeJsonBody(r *http.Request, obj interface{}) error {
	var jsonDecoder = json.NewDecoder(r.Body)
	return jsonDecoder.Decode(obj)
}