package authentication

import "identity-web-api/model"

type Authentication interface {
	GenerateToken(user *model.User) (string, error)
}
