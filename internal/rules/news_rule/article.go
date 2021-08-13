package news_rule

import "time"

type ArticleGetRequest struct {
	ID uint `uri:"id" binding:"required,gte=1"`
}

type ArticleGetResponse struct {
	ID            uint   `gorm:"primarykey"`
	Title         string `json:"title"`
	Desc          string `json:"desc"`
	CoverImageUrl string `json:"cover_image_url"`
	Content       string `json:"content"`
	CreatedBy     string `json:"created_by"`
	UpdatedBy     string `json:"updated_by"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type ArticleListRequest struct {
}

type ArticleCreateRequest struct {
	Title         string `form:"title" json:"title" binding:"required,min=2,max=100"`
	Desc          string `form:"desc" json:"desc" binding:"required,min=2,max=255"`
	CoverImageUrl string `form:"cover_image_url" json:"cover_image_url" binding:"omitempty,url"`
	Content       string `form:"content" json:"content" binding:"required,min=2,max=4294967295"`
	CreatedBy     string `form:"created_by" json:"created_by" binding:"required,min=2,max=100"`
	UpdatedBy     string `form:"updated_by" json:"updated_by" binding:"required,min=2,max=100"`
}

type ArticleUpdateRequest struct {
	ID            uint    `uri:"id" binding:"required,gte=1"`
	Title         *string `form:"title" json:"title" binding:"omitempty,min=2,max=100"`
	Desc          *string `form:"desc" json:"desc" binding:"omitempty,min=2,max=255"`
	CoverImageUrl *string `form:"cover_image_url" json:"cover_image_url" binding:"omitempty,url"`
	Content       *string `form:"content" json:"content" binding:"omitempty,min=2,max=4294967295"`
	Status        *uint8  `form:"status" json:"status" binding:"omitempty,oneof=0 1"`
	UpdatedBy     *string `form:"updated_by" json:"updated_by" binding:"omitempty,min=2,max=100"`
}

type ArticleDeleteRequest struct {
	ID uint `uri:"id" binding:"required,gte=1"`
}
