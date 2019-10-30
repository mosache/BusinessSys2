package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func init() {
	register("user", func(e *gin.Engine) {
		api := e.Group("/api")
		{
			api.GET("/login", login)
		}
	})
}

func login(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"msg":    "success",
	})
}
