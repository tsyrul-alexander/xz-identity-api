package response

import (
	"github.com/google/uuid"
	"github.com/tsyrul-alexander/identity-web-api/model"
)

type UserInfo struct {
	Id uuid.UUID
	Roles []model.UserRole
}

func CreateUserInfo(id uuid.UUID, roles []model.UserRole) *UserInfo {
	return &UserInfo{Id:id, Roles:roles}
}
