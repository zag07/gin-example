package errcode

var (
	Success                   = NewError(0, "成功")
	ServerError               = NewError(10000001, "服务内部错误")
	InvalidParams             = NewError(10000002, "入参错误")
	NotFound                  = NewError(10000003, "找不到")
	UnauthorizedTokenNotFound = NewError(10000004, "校验失败，未携带token")
	UnauthorizedAuthNotExist  = NewError(10000005, "校验失败，用户名或密码错误")
	UnauthorizedTokenError    = NewError(10000006, "校验失败，Token 错误")
	UnauthorizedTokenTimeout  = NewError(10000007, "校验失败，Token 超时")
	UnauthorizedTokenGenerate = NewError(10000008, "校验失败，Token 生成失败")
	TooManyRequests           = NewError(10000009, "请求过多")
)
