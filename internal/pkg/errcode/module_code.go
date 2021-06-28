package errcode

var (
	ArticleGetFail    = NewError(20010001, "文章获取失败")
	ArticleCreateFail = NewError(20010003, "文章创建失败")
	ArticleUpdateFail = NewError(20010004, "文章更新失败")
	ArticleDeleteFail = NewError(20010005, "文章删除失败")
)
