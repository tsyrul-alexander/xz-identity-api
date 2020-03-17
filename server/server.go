package server

import (
	"github.com/gorilla/mux"
	"identity-web-api/controller"
	"identity-web-api/storage"
	"net/http"
	"strconv"
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
	var authorizationController = controller.AuthorizationController{Storage: s.Storage}
	var authenticationController = controller.AuthenticationController{Storage: s.Storage}
	router.HandleFunc("/authorization/register", authorizationController.Registration)
	router.HandleFunc("/", authenticationController.Hello)
	return router
}
