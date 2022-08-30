package models

import (
	"crypto/sha1"
	"encoding/hex"
)

type User struct {
	Id           string `json:"id" bson:"_id,omitempty"`
	Username     string `json:"username" bson:"username"`
	PasswordHash string `json:"-" bson:"password"`
}

func NewUser(userDto UserDto) *User {
	hashPassword := hashPassword(userDto.Password)
	return &User{
		Username:     userDto.Username,
		PasswordHash: hashPassword,
	}
}

func hashPassword(raw string) string {
	h := sha1.New()
	h.Write([]byte(raw))
	return hex.EncodeToString(h.Sum(nil))
}

type UserDto struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required" `
}
