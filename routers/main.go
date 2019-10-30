package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func init() {
	register("main", func(e *gin.Engine) {
		api := e.Group("/api")
		{
			api.POST("/main", mainPage)
		}
	})

}

func mainPage(c *gin.Context) {
	c.JSON(http.StatusOK, ResponseOk(gin.H{
		"status": 200,
		"msg":    "main page success",
	}))
}
