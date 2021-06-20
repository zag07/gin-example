package routes

import "github.com/gin-gonic/gin"

func NewRouter() *gin.Engine {
	r := gin.New()

	setApiRouter(r)

	return r
}
