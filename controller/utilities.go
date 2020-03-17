package controller

import (
	"encoding/json"
	"net/http"
)

func setResponse(w http.ResponseWriter, data interface{}) error {
	w.Header().Add("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(data)
}

func decodeJsonBody(r *http.Request, obj interface{}) error {
	var jsonDecoder = json.NewDecoder(r.Body)
	return jsonDecoder.Decode(obj)
}
