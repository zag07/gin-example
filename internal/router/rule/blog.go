package rule

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

type TagGetRequest struct {
	ID uint `uri:"id" binding:"required,numeric"`
}

type TagGetResponse struct {
	ID        uint   `gorm:"primarykey"`
	Name      string `json:"name"`
	Status    uint8  `json:"status"`
	CreatedBy string `json:"created_by"`
	UpdatedBy string `json:"updated_by"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type TagCreateRequest struct {
	Name      string `form:"title" json:"title" binding:"required,min=2,max=100"`
	CreatedBy string `form:"created_by" json:"created_by" binding:"required,min=2,max=100"`
	UpdatedBy string `form:"updated_by" json:"updated_by" binding:"required,min=2,max=100"`
}

type TagUpdateRequest struct {
	ID        uint    `uri:"id" binding:"required,gte=1"`
	Name      *string `form:"title" json:"title" binding:"omitempty,min=2,max=100"`
	Status    *uint8  `form:"status" json:"status" binding:"omitempty,oneof=0 1"`
	CreatedBy *string `form:"created_by" json:"created_by" binding:"omitempty,min=2,max=100"`
	UpdatedBy *string `form:"updated_by" json:"updated_by" binding:"omitempty,min=2,max=100"`
}

type TagDeleteRequest struct {
	ID uint `uri:"id" binding:"required,gte=1"`
}

type UploadFileRequest struct {
	Type uint   `form:"type" binding:"required"`
	File string `file:"file" binding:"required"`
}

type UploadFileResponse struct {
	FileUrl string `json:"file_url"`
}

type LoginRequest struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

type LoginResponse struct {
	Token string `json:"token"`
}
