package routes

import "github.com/gin-gonic/gin"

func setApiRouter(r *gin.Engine) {
	apiV1 := r.Group("/api/v1")
	{
		apiV1.GET("/article", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"errcode": 0,
				"errmsg":  "ok",
				"data":    "哈哈",
			})
		})
	}
}
