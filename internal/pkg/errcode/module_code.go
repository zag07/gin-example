package errcode

var (
	Err = NewError(20000000, "错误")

	GetArticle    = NewError(20010001, "获取文章失败")
	ListArticle   = NewError(20010002, "获取文章列表失败")
	CreateArticle = NewError(20010003, "创建文章失败")
	UpdateArticle = NewError(20010004, "更新文章失败")
	DeleteArticle = NewError(20010005, "删除文章失败")

	GetTag    = NewError(20010101, "获取标签失败")
	ListTag   = NewError(20010102, "获取标签列表失败")
	CreateTag = NewError(20010103, "创建标签失败")
	UpdateTag = NewError(20010104, "更新标签失败")
	DeleteTag = NewError(20010105, "删除标签失败")

	UploadFileFail = NewError(70000001, "文件上传失败")
)
