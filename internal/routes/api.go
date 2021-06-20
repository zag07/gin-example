package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/zs368/gin-example/internal/app/controllers"
)

func setApiRouter(r *gin.Engine) {
	apiV1 := r.Group("/api/v1")
	{
		user := controllers.NewUser()
		apiV1.GET("/user/:id", user.Get)
	}
}
