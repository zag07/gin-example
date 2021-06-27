package models

type ArticleTag struct {
	ID        uint   `gorm:"primaryKey"`
	ArticleID uint32 `json:"article_id"`
	TagID     uint32 `json:"tag_id"`
	State     uint8  `json:"state"`
	CreatedAt uint32 `json:"created_at"`
	CreatedBy string `json:"created_by"`
	UpdatedAt string `json:"updated_at"`
	UpdatedBy uint32 `json:"updated_by"`
}

func (ArticleTag) TableName() string {
	return "article_tag"
}
