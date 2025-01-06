package userService

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email        string `json:"email"`
	PasswordHash string `json:"password_hash"`
}
