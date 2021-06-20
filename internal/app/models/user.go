package models

type User struct {
	*Model
	Name            string `json:"name"`
	Email           string `json:"email"`
	EmailVerifiedAt string `json:"email_verified_at"`
	Password        string `json:"password"`
	RememberToken   uint8  `json:"remember_token"`
}

func (User) TableName() string {
	return "user"
}
