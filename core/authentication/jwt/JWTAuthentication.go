package jwt

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/tsyrul-alexander/identity-web-api/model/data"
)

var NotValidTokenError = errors.New("not valid token")

type Authentication struct {
	JwtKey string
}

func (jwtAuth *Authentication) GenerateToken(user *data.User) (string, error) {
	var token = jwt.NewWithClaims(jwt.SigningMethodHS256, createClaims(user))
	return token.SignedString([]byte(jwtAuth.JwtKey))
}

func createClaims(user *data.User) *Claim {
	return &Claim {
		UserId: user.ID,
	}
}

func (jwtAuth *Authentication) GetUserId(token string) (uuid.UUID, error) {
	var claim = &Claim{}
	t, err := jwt.ParseWithClaims(token, claim, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtAuth.JwtKey), nil
	})
	if err != nil {
		return uuid.Nil, nil
	}
	if !t.Valid {
		return uuid.Nil, NotValidTokenError
	}
	return claim.UserId, nil
}

