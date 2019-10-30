package middleware

import (
	routes "BusinessSys/routers"
	"github.com/gin-gonic/gin"
	"net/http"
)

func init() {
	registerWithWeight(1, jwtMiddleWare)
}

func jwtMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		//Get请求和login请求不验证token
		if c.Request.Method == http.MethodGet || c.Request.URL.String() == "/api/login" {
			c.Next()
			return
		}
		token := c.PostForm("token")

		if len(token) == 0 {
			c.JSON(http.StatusBadRequest, routes.ResponseBad("token is not exist"))
			c.Abort()
			return
		}

		c.Next()
	}
}
