package controller

import (
	"encoding/json"
	"log"
	"net/http"
)

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
	SetResponse(w, data)
}

func decodeJsonBody(r *http.Request, obj interface{}) error {
	var jsonDecoder = json.NewDecoder(r.Body)
	return jsonDecoder.Decode(obj)
}
