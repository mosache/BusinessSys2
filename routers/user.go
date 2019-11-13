package routes

import (
	"BusinessSys/model"
	"BusinessSys/service"
	"BusinessSys/utils"
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

// @Summary 登录接口
// @Description 登录
// @Accept  multipart/form-data
// @Produce json
// @Param username formData string true "用户名"
// @Param password formData string false "密码"
// @Router /api/login [post]
func login(c *gin.Context) {
	var (
		loginInUser *model.User
		isSuccess   bool
	)
	username := c.PostForm("username")
	pwd := c.PostForm("password")

	if len(username) == 0 {
		ServiceFailJson("用户名不能为空", c)
		return
	}

	if len(pwd) == 0 {
		ServiceFailJson("密码不能为空", c)
		return
	}

	if loginInUser, isSuccess = service.LoginIn(username, pwd); !isSuccess {
		ServiceFailJson("用户名或密码错误", c)
		return
	}

	token := utils.GetNewToken(loginInUser.UserID)

	ServiceOKJson(gin.H{
		"token": token,
	}, c)
}


