package router

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/zs368/gin-example/internal/pkg/chatroom"
	"github.com/zs368/gin-example/internal/service"
)

var e service.ExampleService

func SetApiRouter(r *gin.Engine) {
	c := r.Group("/c")
	{
		c.POST("/upload/file", e.UploadFile)
		c.StaticFS("/static", http.Dir(""))
	}

	apiV1 := r.Group("/api/v1")
	// Use(middleware.JWT()).
	// Use(middleware.Translations())
	// Use(middleware.LoggerWithZap()).
	{
		apiV1.GET("/article/:id", e.GetArticle)
		// apiV1.GET("/article/page/:id", e.ListArticle)
		apiV1.POST("/article", e.CreateArticle)
		apiV1.PUT("/article/:id", e.UpdateArticle)
		apiV1.DELETE("/article/:id", e.DeleteArticle)

		/*apiV1.GET("/tag/:id", e.GetTag)
		apiV1.GET("/tag/page/:id", e.ListTag)
		apiV1.POST("/tag", e.CreateTag)
		apiV1.PUT("/tag/:id", e.UpdateTag)
		apiV1.DELETE("/tag/:id", e.DeleteTag)*/
	}
}

func SetWSRouter(r *gin.Engine) {
	go chat_svs.Broadcaster.Run()

	c := r.Group("/chat")
	{
		c.GET("/", e.Home)
		c.GET("/2", e.Home2)
		c.GET("/user_list", e.UserList)
		c.GET("/ws", e.WS)
	}
}
