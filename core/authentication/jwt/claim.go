package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

type Claim struct {
	UserId uuid.UUID
	jwt.StandardClaims
}
