package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
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
	var validate UserGet
	if err := c.ShouldBindUri(&validate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, "aaa")
}
