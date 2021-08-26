package router

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	_ "github.com/zs368/gin-example/docs"
	"github.com/zs368/gin-example/internal/conf"
)

func NewRouter(cfg *conf.HTTP) *gin.Engine {
	r := gin.New()

	if cfg.Debug == true {
		gin.SetMode("debug")
	} else {
		gin.SetMode("release")
	}


	r.Use(gin.Logger(), gin.Recovery())

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	SetApiRouter(r)

	SetWSRouter(r)

	return r
}
