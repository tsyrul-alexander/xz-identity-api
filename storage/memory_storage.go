package storage

import (
	"github.com/google/uuid"
	"github.com/tsyrul-alexander/xz-identity-api/model/memory"
)

type MemoryStorage interface {
	SetUser(user *memory.User) bool
	GetUser(id uuid.UUID) (*memory.User, bool)
}
