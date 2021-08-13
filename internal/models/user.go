package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name            string `json:"name"`
	Email           string `json:"email"`
	EmailVerifiedAt string `json:"email_verified_at"`
	Password        string `json:"password"`
	RememberToken   uint8  `json:"remember_token"`
	Role            string `json:"role"`
}

func (User) TableName() string {
	return "user"
}
