package routes

import (
	"BusinessSys/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v8"
)

func init() {
	register("notice", func(e *gin.Engine) {
		api := e.Group("/api")
		{
			api.POST("/addNoticePoint",AddNoticePoint)

		}

	})
}


// @summary 添加印象点
// @Param name formData string true "名称"
// @Param type formData int true "类型"
// @Param coordinateX formData string true "纬度"
// @Param coordinateY formData string true "经度"
// @Param feel_now formData int false "当前印象"
// @Router /api/addNoticePoint [post]
func AddNoticePoint(c *gin.Context){
	var (
		noticePoint *model.NoticePoint
	)

	noticePoint = &model.NoticePoint{}

	if err := c.ShouldBind(noticePoint);err!=nil{
		ServiceFailJson(noticePoint.GetError(err.(validator.ValidationErrors)),c)
		return
	}

	fmt.Println(noticePoint)

	ServiceOKJson(nil,c)
}