// https://github.com/gin-gonic/gin 撸一下文档就行

package test

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Article struct {
	Id   string  `json:"id"`
	Name *string `json:"name,omitempty"`
	Desc *string `json:"desc,omitempty"`
}

func Test_JSON_Empty(t *testing.T) {
	jsonData := `{"id":"1234","name":"xyz","desc":""}`
	req := Article{}
	_ = json.Unmarshal([]byte(jsonData), &req)
	fmt.Printf("%+v\n", req)
	fmt.Printf("%s\n", *req.Name)
	fmt.Printf("%s\n", *req.Desc)
}

func Test_JSON_Nil(t *testing.T) {
	jsonData := `{"id":"1234","name":"xyz"}`
	req := Article{}
	_ = json.Unmarshal([]byte(jsonData), &req)
	fmt.Printf("%+v\n", req)
	fmt.Printf("%s\n", *req.Name)
	fmt.Printf("%+v\n", req.Desc == nil)
}
