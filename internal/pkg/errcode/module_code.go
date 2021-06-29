package errcode

var (
	ArticleGetFail    = NewError(20010001, "文章获取失败")
	ArticleCreateFail = NewError(20010003, "文章创建失败")
	ArticleUpdateFail = NewError(20010004, "文章更新失败")
	ArticleDeleteFail = NewError(20010005, "文章删除失败")

	TagGetFail    = NewError(20010101, "标签获取失败")
	TagCreateFail = NewError(20010103, "标签创建失败")
	TagUpdateFail = NewError(20010104, "标签更新失败")
	TagDeleteFail = NewError(20010105, "标签删除失败")
)
