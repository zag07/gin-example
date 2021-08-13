package models

import "gorm.io/gorm"

type ArticleTag struct {
	gorm.Model
	ArticleID uint32 `json:"article_id"`
	TagID     uint32 `json:"tag_id"`
	Status    uint8  `json:"status"`
	CreatedBy string `json:"created_by"`
	UpdatedBy uint32 `json:"updated_by"`
}

func (ArticleTag) TableName() string {
	return "article_tag"
}
