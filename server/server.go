package server

import (
	"../controller"
	"../storage"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type Server struct {
	Config Config
	Storage storage.Storage
}

func Create(config Config, storage storage.Storage) *Server {
	return &Server{Config:config, Storage:storage}
}

func (s *Server) Start() error {
	var router = s.UseRouting()
	var serverAddress = s.Config.Ip + ":" + strconv.Itoa(s.Config.Port)
	return http.ListenAndServe(serverAddress, router)
}

func (s *Server) UseRouting() *mux.Router {
	var router = mux.NewRouter()
	router.HandleFunc("/authorization/register", controller.Registration)
	return router
}