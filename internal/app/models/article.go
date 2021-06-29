package models

import "gorm.io/gorm"

type Article struct {
	ID            uint32 `gorm:"primaryKey"`
	Title         string `json:"title"`
	Desc          string `json:"desc"`
	CoverImageUrl string `json:"cover_image_url"`
	Content       string `json:"content"`
	State         uint8  `json:"state"`
	CreatedAt     uint32 `json:"created_at"`
	CreatedBy     string `json:"created_by"`
	UpdatedAt     uint32 `json:"updated_at"`
	UpdatedBy     string `json:"updated_by"`
}

func (Article) TableName() string {
	return "article"
}

func (a Article) Create(db *gorm.DB) (*Article, error) {
	if err := db.Create(&a).Error; err != nil {
		return nil, err
	}

	return &a, nil
}

func (a Article) Update(db *gorm.DB, values interface{}) error {
	if err := db.Model(&a).Where("id = ? AND is_del = ?", a.ID, 0).Updates(values).Error; err != nil {
		return err
	}

	return nil
}

func (a Article) Delete(db *gorm.DB) error {
	if err := db.Where("id = ? AND is_del = ?", a.ID, 0).Delete(&a).Error; err != nil {
		return err
	}

	return nil
}
