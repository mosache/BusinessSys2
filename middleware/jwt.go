package middleware

import (
	routes "BusinessSys/routers"
	"BusinessSys/service"
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
			routes.ServiceFailJson("token 不能为空", c)
			c.Abort()
			return
		}

		data, err := service.CheckToken(token)

		if err != nil {
			routes.ServiceFailJson("token 解析失败", c)
			c.Abort()
			return
		}

		var tokenData *service.TokenData
		var isOk bool

		if tokenData, isOk = data.(*service.TokenData); !isOk {
			routes.ServiceFailJson("token data 解析失败", c)
			c.Abort()
			return
		}

		if tokenData.UserId != 1 {
			routes.ServiceFailJson("token data 解析失败", c)
			c.Abort()
			return
		}
		c.Next()
	}
}
