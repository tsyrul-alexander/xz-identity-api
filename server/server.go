package server

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/tsyrul-alexander/identity-web-api/controller"
	"github.com/tsyrul-alexander/identity-web-api/core/authentication"
	"github.com/tsyrul-alexander/identity-web-api/core/authentication/jwt"
	"github.com/tsyrul-alexander/identity-web-api/setting"
	"github.com/tsyrul-alexander/identity-web-api/storage"
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
	router.Use(handlers.CORS(handlers.AllowedOrigins([]string{"*"})))
	var authorizationController = controller.AuthorizationController{Storage: s.Storage, Authentication: getAuthenticationMethod()}
	var authenticationController = controller.AuthenticationController{Storage: s.Storage, Authentication: getAuthenticationMethod()}
	router.HandleFunc("/authorization/register", authorizationController.Registration).Methods("POST")
	router.HandleFunc("/authorization/login", authenticationController.Login).Methods("POST")
	router.HandleFunc("/ping", func(writer http.ResponseWriter, request *http.Request) {
		controller.SetResponse(writer, "pong")
	})
	return router
}

func getAuthenticationMethod() authentication.Authentication {
	var s = setting.GetAppSetting()
	return &jwt.Authentication{JwtKey: s.JwtKey}
}
