package rules

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

type ArticleCreateRequest struct {
	Title         string `json:"title" binding:"required,min=2,max=100"`
	Desc          string `json:"desc" binding:"required,min=2,max=255"`
	CoverImageUrl string `json:"cover_image_url" binding:"omitempty,url"`
	Content       string `json:"content" binding:"required,min=2,max=4294967295"`
	CreatedBy     string `json:"created_by" binding:"required,min=2,max=100"`
	UpdatedBy     string `json:"updated_by" binding:"required,min=2,max=100"`
}

type ArticleUpdateRequest struct {
	ID            uint   `json:"id" binding:"required,gte=1"`
	Title         string `json:"title" binding:"omitempty,min=2,max=100"`
	Desc          string `json:"desc" binding:"omitempty,min=2,max=255"`
	CoverImageUrl string `json:"cover_image_url" binding:"omitempty,url"`
	Content       string `json:"content" binding:"omitempty,min=2,max=4294967295"`
	State         uint8  `json:"state" binding:"omitempty,oneof=0 1"`
	UpdatedBy     string `json:"updated_by" binding:"omitempty,min=2,max=100"`
}

type ArticleDeleteRequest struct {
	ID uint `uri:"id" binding:"required,gte=1"`
}