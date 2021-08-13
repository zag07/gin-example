package models

import "gorm.io/gorm"

type Tag struct {
	gorm.Model
	Name      string `json:"name"`
	Status    uint8  `json:"status"`
	CreatedBy string `json:"created_by"`
	UpdatedBy string `json:"updated_by"`
}

func (Tag) TableName() string {
	return "tag"
}
