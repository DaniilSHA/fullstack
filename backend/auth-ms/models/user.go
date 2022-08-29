package models

type User struct {
	Id           int    `json:"id" bson:"_id,omitempty"`
	Username     string `json:"username" bson:"username"`
	PasswordHash string `json:"-" bson:"password"`
}

type UserDto struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required" `
}
