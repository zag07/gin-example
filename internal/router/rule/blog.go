package rule

type GetArticle struct {
	ID int64 `uri:"id" binding:"required,gte=1"`
}

type ListArticle struct {
}

type CreateArticle struct {
	Title         string `form:"title" json:"title" binding:"required,min=2,max=100"`
	Desc          string `form:"desc" json:"desc" binding:"required,min=2,max=255"`
	CoverImageUrl string `form:"cover_image_url" json:"cover_image_url" binding:"omitempty,url"`
	Content       string `form:"content" json:"content" binding:"required,min=2,max=4294967295"`
	CreatedBy     string `form:"created_by" json:"created_by" binding:"required,min=2,max=100"`
	UpdatedBy     string `form:"updated_by" json:"updated_by" binding:"required,min=2,max=100"`
}

// UpdateArticle 使用指针来判断是否传入
type UpdateArticle struct {
	ID            int64   `uri:"id" binding:"required,gte=1"`
	Title         *string `form:"title" json:"title" binding:"omitempty,min=2,max=100"`
	Desc          *string `form:"desc" json:"desc" binding:"omitempty,min=2,max=255"`
	CoverImageUrl *string `form:"cover_image_url" json:"cover_image_url" binding:"omitempty,url"`
	Content       *string `form:"content" json:"content" binding:"omitempty,min=2,max=4294967295"`
	Status        *int8   `form:"status" json:"status" binding:"omitempty,oneof=0 1"`
	UpdatedBy     *string `form:"updated_by" json:"updated_by" binding:"omitempty,min=2,max=100"`
}

type DeleteArticle struct {
	ID int64 `uri:"id" binding:"required,gte=1"`
}

type GetTag struct {
	ID int64 `uri:"id" binding:"required,numeric"`
}

type CreateTag struct {
	Name      string `form:"title" json:"title" binding:"required,min=2,max=100"`
	CreatedBy string `form:"created_by" json:"created_by" binding:"required,min=2,max=100"`
	UpdatedBy string `form:"updated_by" json:"updated_by" binding:"required,min=2,max=100"`
}

// UpdateTag 使用参数校验来判断是否传入（即：全部传入）
type UpdateTag struct {
	ID        int64  `uri:"id" binding:"required,gte=1"`
	Name      string `form:"title" json:"title" binding:"required,min=2,max=100"`
	Status    int8  `form:"status" json:"status" binding:"required,oneof=0 1"`
	CreatedBy string `form:"created_by" json:"created_by" binding:"required,min=2,max=100"`
	UpdatedBy string `form:"updated_by" json:"updated_by" binding:"required,min=2,max=100"`
}

type DeleteTag struct {
	ID int64 `uri:"id" binding:"required,gte=1"`
}

type UploadFile struct {
	Type uint   `form:"type" binding:"required"`
	File string `file:"file" binding:"required"`
}

type Login struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}
