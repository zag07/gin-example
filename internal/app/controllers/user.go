package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/zs368/gin-example/internal/app/models"
	"github.com/zs368/gin-example/internal/pkg/app"
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

	var (
		r   = app.NewResponse(c)
		err error
	)
	if err = c.ShouldBindUri(&params); err != nil {
		r.ToErrorResponse(errcode.InvalidParams.WithDetails(err.Error()))
		return
	}

	var (
		db   = database.DB
		user models.User
	)
	db.Where("id = ?", params.ID).First(&user)

	r.ToResponse(user)
}
