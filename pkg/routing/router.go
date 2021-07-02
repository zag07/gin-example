package routing

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	_ "github.com/zs368/gin-example/docs"
	"github.com/zs368/gin-example/internal/app/middleware"
	"github.com/zs368/gin-example/internal/routes"
)

func NewRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger(), gin.Recovery())

	r.Use(middleware.Translations())

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	routes.SetApiRouter(r)

	return r
}
