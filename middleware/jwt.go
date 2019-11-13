package middleware

import (
	"BusinessSys/model"
	routes "BusinessSys/routers"
	"BusinessSys/service"
	"BusinessSys/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func init() {
	registerWithWeight(1, jwtMiddleWare)
}

func jwtMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		//Get请求和login请求不验证token
		if c.Request.Method == http.MethodGet || c.Request.URL.String() == "/api/login" || strings.HasPrefix(c.Request.URL.String(), "/swagger/") {
			c.Next()
			return
		}
		token := c.PostForm("token")

		if len(token) == 0 {
			routes.ServiceFailJson("token 不能为空", c)
			c.Abort()
			return
		}

		data, err := utils.CheckToken(token)

		if err != nil {

			fmt.Println(err.Error())

			routes.ServiceFailJson("token 解析失败", c)
			c.Abort()
			return
		}

		var tokenData *utils.TokenData
		var isOk bool

		if tokenData, isOk = data.(*utils.TokenData); !isOk {
			routes.ServiceFailJson("token data 解析失败", c)
			c.Abort()
			return
		}

		//token通过验证
		var currentUser *model.User
		if currentUser,err = service.GetUserByID(tokenData.UserId);err!=nil {
			routes.ServiceFailJson("token data 解析失败", c)
		}

		c.Set("currentUser",currentUser)


		//if tokenData.UserId != 1 {
		//	routes.ServiceFailJson("token data 解析失败", c)
		//	c.Abort()
		//	return
		//}
		c.Next()
	}
}
