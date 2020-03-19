package model

import (
	"golang.org/x/crypto/bcrypt"
)

type HashPassword string

func (hashedPassword *HashPassword) String() string {
	return string(*hashedPassword)
}

func (hashedPassword *HashPassword) GetIsCompareHashPassword(password string) bool  {
	var hashedPasswordBytes = []byte(hashedPassword.String())
	var passwordBytes = []byte(password)
	if  err := bcrypt.CompareHashAndPassword(hashedPasswordBytes, passwordBytes); err != nil {
		return false
	}
	return true
}

func CreateHashPassword(password string) HashPassword  {
	var bytes, _ = bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	var value = string(bytes)
	return HashPassword(value)
}
