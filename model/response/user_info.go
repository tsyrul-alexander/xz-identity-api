package response

import (
	"github.com/google/uuid"
	"github.com/tsyrul-alexander/xz-identity-api/model"
)

type UserInfo struct {
	Id uuid.UUID
	Roles []model.UserRole
}

func CreateUserInfo(id uuid.UUID, roles []model.UserRole) *UserInfo {
	return &UserInfo{Id:id, Roles:roles}
}

func (ui *UserInfo) GetIfExistRole(role model.UserRole) bool {
	for _, r := range ui.Roles {
		if int(r) == int(role) {
			return true
		}
	}
	return false
}