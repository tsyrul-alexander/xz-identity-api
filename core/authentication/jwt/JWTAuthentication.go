package jwt

import (
	jwt "github.com/dgrijalva/jwt-go"
	"identity-web-api/model"
)

type Authentication struct {
	JwtKey string
}

func (jwtAuth *Authentication) GenerateToken(user *model.User) (string, error)  {
	var token = jwt.NewWithClaims(jwt.SigningMethodHS256, createClaims(user))
	return token.SignedString([]byte(jwtAuth.JwtKey))
}

func createClaims(user *model.User) *jwt.MapClaims {
	return &jwt.MapClaims {
		"user_id": user.ID.String(),
	}
}

