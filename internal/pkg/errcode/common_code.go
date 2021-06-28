package errcode

var (
	Success                   = NewError(0, "成功")
	ServerError               = NewError(10000001, "服务内部错误")
	InvalidParams             = NewError(10000002, "入参错误")
	NotFound                  = NewError(10000003, "找不到")
	UnauthorizedAuthNotExist  = NewError(10000004, "鉴权失败，找不到对应的 AppKey 和 AppSecret")
	UnauthorizedTokenError    = NewError(10000005, "鉴权失败，Token 错误")
	UnauthorizedTokenTimeout  = NewError(10000006, "鉴权失败，Token 超时")
	UnauthorizedTokenGenerate = NewError(10000007, "鉴权失败，Token 生成失败")
	TooManyRequests           = NewError(10000008, "请求过多")
)
