package controller

import "net/http"

type ServiceError struct {
	Message string
	Code int
	StatusCode int
}
var (
	DbError                = &ServiceError{"db error", 100, http.StatusInternalServerError}
	InvalidRequest         = &ServiceError{"Invalid request", 101,http.StatusBadRequest}
	UserExistsError        = &ServiceError{"User exists", 102, http.StatusInternalServerError}
	AuthenticationRequired = &ServiceError{"Invalid request", 104,http.StatusNetworkAuthenticationRequired}
	InvalidCredential      = &ServiceError{"Invalid credentials", 105,http.StatusInternalServerError}
	GenerateTokenError     = &ServiceError{"Generate token error", 106,http.StatusInternalServerError}
	ParseTokenError        = &ServiceError{"Get token error", 107,http.StatusInternalServerError}
)
