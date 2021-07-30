package models

import "gorm.io/gorm"

type Article struct {
	gorm.Model
	Title         string `json:"title"`
	Desc          string `json:"desc" gorm:"default:haha"`
	CoverImageUrl string `json:"cover_image_url"`
	Content       string `json:"content"`
	Status        uint8  `json:"status"`
	CreatedBy     string `json:"created_by"`
	UpdatedBy     string `json:"updated_by"`
}

func (Article) TableName() string {
	return "article"
}

func (a *Article) BeforeCreate(tx *gorm.DB) error {
	a.CoverImageUrl = "https://google.com/"
	return nil
}
