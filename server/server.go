package server

import (
	"IdentityWebApi/controller"
	"IdentityWebApi/storage"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//Server ...
type Server struct {
	Config  Config
	Storage storage.Storage
}

//Create ...
func Create(config Config, storage storage.Storage) *Server {
	return &Server{Config: config, Storage: storage}
}

//Start ...
func (s *Server) Start() error {
	var router = s.UseRouting()
	var serverAddress = s.Config.IP + ":" + strconv.Itoa(s.Config.Port)
	return http.ListenAndServe(serverAddress, router)
}

//UseRouting ...
func (s *Server) UseRouting() *mux.Router {
	var router = mux.NewRouter()
	var authorizationController = controller.AuthorizationController{Server: s}
	router.HandleFunc("/authorization/register", authorizationController.Registration)
	return router
}
