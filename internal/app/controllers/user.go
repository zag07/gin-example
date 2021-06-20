package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zs368/gin-example/internal/app/models"
)

type User struct {
}

func NewUser() User {
	return User{}
}

type UserGet struct {
	ID string `uri:"id" binding:"required,numeric"`
}

func (u User) Get(c *gin.Context) {
	var userGet UserGet
	if err := c.ShouldBindUri(&userGet); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db, err := models.NewDBEngine()
	if err != nil {
		c.Error(err)
		return
	}

	var user []models.User
	db.Where("id = ?", userGet.ID).Find(&user)

	c.JSON(200, user)
}
