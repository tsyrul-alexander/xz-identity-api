package model

import "github.com/google/uuid"

type DefaultIdentity struct {
	ID uuid.UUID
	Login string
	Password HashPassword
}
