package routing

import (
	"github.com/gin-gonic/gin"
	"github.com/zs368/gin-example/internal/routes"
)

func NewRouter() *gin.Engine {
	r := gin.New()

	routes.SetApiRouter(r)

	return r
}
