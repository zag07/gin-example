package test

import (
	"bytes"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin/binding"
	"github.com/stretchr/testify/assert"
)

func TestValidationSuccess(t *testing.T) {
	type HogeStruct struct {
		Hoge *int `json:"hoge" binding:"exists"`
	}

	var obj HogeStruct
	req := requestWithBody("POST", "/", `{}`)
	err := binding.JSON.Bind(req, &obj)
	assert.NoError(t, err)
	assert.Equal(t, "FOO", obj.Hoge)
}

func requestWithBody(method, path, body string) (req *http.Request) {
	req, _ = http.NewRequest(method, path, bytes.NewBufferString(body))
	return
}