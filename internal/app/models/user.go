package models

type User struct {
	ID              uint   `gorm:"primaryKey"`
	Name            string `json:"name"`
	Email           string `json:"email"`
	EmailVerifiedAt string `json:"email_verified_at"`
	Password        string `json:"password"`
	RememberToken   uint8  `json:"remember_token"`
	CreatedAt       uint32 `json:"created_at"`
	UpdatedAt       string `json:"updated_at"`
}

func (User) TableName() string {
	return "user"
}
