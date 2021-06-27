package models

type Tag struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `json:"name"`
	State     uint8  `json:"state"`
	CreatedAt uint32 `json:"created_at"`
	CreatedBy string `json:"created_by"`
	UpdatedAt string `json:"updated_at"`
	UpdatedBy uint32 `json:"updated_by"`
}

func (Tag) TableName() string {
	return "tag"
}
