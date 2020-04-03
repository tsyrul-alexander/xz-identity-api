package server

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/tsyrul-alexander/xz-identity-api/controller"
	"github.com/tsyrul-alexander/xz-identity-api/core/authentication"
	"github.com/tsyrul-alexander/xz-identity-api/core/authentication/jwt"
	"github.com/tsyrul-alexander/xz-identity-api/setting"
	"github.com/tsyrul-alexander/xz-identity-api/storage"
	"net/http"
	"strconv"
)

//Server ...
type Server struct {
	Config        Config
	DataStorage   storage.DataStorage
	MemoryStorage storage.MemoryStorage
}

//Create ...
func Create(config Config, storage storage.DataStorage, memoryStorage storage.MemoryStorage) *Server {
	return &Server{Config: config, DataStorage: storage, MemoryStorage: memoryStorage}
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
	var authorizationController = s.configureAuthorizationController()
	var authenticationController = s.configureAuthenticationController()
	router.HandleFunc("/authorization/register", authorizationController.Registration).Methods("POST")
	router.HandleFunc("/authorization/login", authenticationController.Login).Methods("GET")
	router.HandleFunc("/authorization/get-user-info", authenticationController.GetUserInfo).Methods("GET")
	router.HandleFunc("/ping", func(writer http.ResponseWriter, request *http.Request) {
		controller.SetResponse(writer, "pong")
	})
	return router
}

func (s *Server) configureAuthorizationController() *controller.AuthorizationController {
	return &controller.AuthorizationController{Storage: s.DataStorage, Authentication: getAuthenticationMethod()}
}

func (s *Server) configureAuthenticationController() *controller.AuthenticationController {
	return &controller.AuthenticationController{
		DataStorage:    s.DataStorage,
		Authentication: getAuthenticationMethod(),
		MemoryStorage:  s.MemoryStorage,
	}
}

func getAuthenticationMethod() authentication.Authentication {
	var s = setting.GetAppSetting()
	return &jwt.Authentication{JwtKey: s.Authorized.Jwt.Key}
}
