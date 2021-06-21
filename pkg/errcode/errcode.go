package errcode

import (
	"fmt"
	"net/http"
)

type Error struct {
	errcode int
	errmsg  string
	details []string
}

var codes = map[int]string{}

func NewError(errcode int, errmsg string) *Error {
	if _, ok := codes[errcode]; ok {
		panic(fmt.Sprintf("错误码 %d 已经存在，请更换一个", errcode))
	}
	codes[errcode] = errmsg
	return &Error{errcode: errcode, errmsg: errmsg}
}

func (e *Error) Error() string {
	return fmt.Sprintf("错误码：%d, 错误信息:：%s", e.Code(), e.Msg())
}

func (e *Error) Code() int {
	return e.errcode
}

func (e *Error) Msg() string {
	return e.errmsg
}

func (e *Error) Details() []string {
	return e.details
}

func (e *Error) StatusCode() int {
	switch e.Code() {
	case Success.Code():
		return http.StatusOK
	case ServerError.Code():
		return http.StatusInternalServerError
	case InvalidParams.Code():
		return http.StatusBadRequest
	case UnauthorizedAuthNotExist.Code():
		fallthrough
	case UnauthorizedTokenError.Code():
		fallthrough
	case UnauthorizedTokenTimeout.Code():
		fallthrough
	case UnauthorizedTokenGenerate.Code():
		return http.StatusUnauthorized
	case NotFound.Code():
		return http.StatusNotFound
	case TooManyRequests.Code():
		return http.StatusTooManyRequests
	}

	return http.StatusInternalServerError
}
