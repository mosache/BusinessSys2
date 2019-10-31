package routes

import (
	"BusinessSys/service"
	"github.com/gin-gonic/gin"
)

func init() {
	register("user", func(e *gin.Engine) {
		api := e.Group("/api")
		{
			api.POST("/login", login)
		}
	})
}

func login(c *gin.Context) {

	token := service.GetNewToken(1)

	ServiceOKJson(gin.H{
		"token": token,
	}, c)
}
