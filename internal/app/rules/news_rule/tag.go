package news_rule

import "time"

type TagGetRequest struct {
	ID uint32 `uri:"id" binding:"required,numeric"`
}

type TagGetResponse struct {
	ID        uint   `gorm:"primarykey"`
	Name      string `json:"name"`
	State     uint8  `json:"state"`
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
	State     *uint8  `form:"state" json:"state" binding:"omitempty,oneof=0 1"`
	CreatedBy *string `form:"created_by" json:"created_by" binding:"omitempty,min=2,max=100"`
	UpdatedBy *string `form:"updated_by" json:"updated_by" binding:"omitempty,min=2,max=100"`
}

type TagDeleteRequest struct {
	ID uint `uri:"id" binding:"required,gte=1"`
}
