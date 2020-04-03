package controller

import (
	"encoding/json"
	"github.com/tsyrul-alexander/xz-identity-api/model/response"
	"log"
	"net/http"
)

const AuthorizedTokenName = "Authorization"

//SetResponse ...
func SetResponse(w http.ResponseWriter, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Println(err.Error())
	}
}

func setError(w http.ResponseWriter, data ServiceError, e error) {
	if e != nil {
		log.Println(e.Error())
	}
	w.WriteHeader(data.StatusCode)
	SetResponse(w, response.Error{Message:data.Message, Code:data.Code})
}

func decodeJsonBody(r *http.Request, obj interface{}) error {
	var jsonDecoder = json.NewDecoder(r.Body)
	return jsonDecoder.Decode(obj)
}

func getAuthorizedToken(r *http.Request) string {
	return r.Header.Get(AuthorizedTokenName)
}
