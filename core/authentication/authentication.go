package authentication

import "github.com/tsyrul-alexander/identity-web-api/model"

type Authentication interface {
	GenerateToken(user *model.User) (string, error)
}
