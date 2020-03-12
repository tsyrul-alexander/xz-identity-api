package server

import "net/http"

type Routing struct {
	RegRxpRule string
	Action func(w http.ResponseWriter, r *http.Request)
}