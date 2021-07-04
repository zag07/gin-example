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
	Name      string `json:"title" binding:"required,min=2,max=100"`
	CreatedBy string `json:"created_by" binding:"required,min=2,max=100"`
	UpdatedBy string `json:"updated_by" binding:"required,min=2,max=100"`
}

type TagUpdateRequest struct {
	ID        uint   `json:"id" binding:"required,gte=1"`
	Name      string `json:"title" binding:"required,min=2,max=100"`
	State     uint8  `json:"state" binding:"oneof=0 1"`
	CreatedBy string `json:"created_by" binding:"required,min=2,max=100"`
	UpdatedBy string `json:"updated_by" binding:"required,min=2,max=100"`
}

type TagDeleteRequest struct {
	ID uint `uri:"id" binding:"required,gte=1"`
}
