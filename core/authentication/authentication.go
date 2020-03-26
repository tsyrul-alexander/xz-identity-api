package authentication

import (
	"github.com/google/uuid"
	"github.com/tsyrul-alexander/identity-web-api/model/data"
)

type Authentication interface {
	GenerateToken(user *data.User) (string, error)
	GetUserId(token string) (uuid.UUID, error)
}
