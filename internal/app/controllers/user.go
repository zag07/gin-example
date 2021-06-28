package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/zs368/gin-example/internal/app/controllers/core"
	"github.com/zs368/gin-example/internal/app/models"
	"github.com/zs368/gin-example/internal/pkg/errcode"
	"github.com/zs368/gin-example/pkg/database"
)

type User struct{}

func NewUser() User {
	return User{}
}

func (u User) Get(c *gin.Context) {
	params := struct {
		ID string `uri:"id" binding:"required,numeric"`
	}{}

	r := core.NewResponse(c)
	if err := c.ShouldBindUri(&params); err != nil {
		r.ToErrorResponse(errcode.InvalidParams.WithDetails(err.Error()))
		return
	}

	var user []models.User

	database.DB.Where("id = ?", params.ID).Find(&user)

	r.ToResponse(user)
}
