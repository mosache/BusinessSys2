package routes

import (
	"BusinessSys/model"
	"BusinessSys/service"
	"BusinessSys/utils"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)

func init() {
	register("project", func(e *gin.Engine) {
		api:= e.Group("/api")
		{
			api.POST("/addProject",InsertProject)
		}
	})
}

// @Summary 新建项目
// @Description 添加项目
// @Param data formData string true "data"
// @Router /api/addProject [post]
func InsertProject(c *gin.Context){
	var (
		jsonData string

	)
	p := &model.Project{}

	jsonData = c.PostForm("data")

	err := json.Unmarshal([]byte(jsonData),p)

	if err !=nil {
		ServiceFailJson(err.Error(),c)
		return
	}

	if err:=checkProject(p);err!=nil{
		ServiceFailJson(err.Error(),c)
		return
	}

	p.AddTime = time.Now()

	user,err := utils.GetUser(c)
	if err != nil{
		ServiceFailJson(err.Error(),c)
		return
	}
	p.AddUserID = user.UserID

	if err := service.InsertProject(p);err!=nil {
		ServiceFailJson(err.Error(),c)
		return
	}

ServiceOKJson(nil,c)
}

func checkProject(p *model.Project) (err error){
	if len(strings.TrimSpace(p.ProjectName)) == 0 {
		err = errors.New("项目名称不能为空")
		return
	}

	if len(strings.TrimSpace(p.Address)) == 0 {
		err = errors.New("项目地址不能为空")
		return
	}

	if len(strings.TrimSpace(p.ProvinceName)) == 0 {
		err = errors.New("省名称不能为空")
		return
	}

	if p.ProvinceID == 0 {
		err = errors.New("省ID不能为空")
		return
	}

	if len(strings.TrimSpace(p.CityName)) == 0 {
		err = errors.New("市名称不能为空")
		return
	}

	if p.CityID == 0 {
		err = errors.New("市ID不能为空")
		return
	}

	if len(strings.TrimSpace(p.AreaName)) == 0 {
		err = errors.New("区域名称不能为空")
		return
	}

	if p.AreaID == 0 {
		err = errors.New("区域ID不能为空")
		return
	}

	if len(strings.TrimSpace(p.CoordinateX)) == 0 {
		err = errors.New("纬度不能为空")
		return
	}

	if len(strings.TrimSpace(p.CoordinateY)) == 0 {
		err = errors.New("经度不能为空")
		return
	}

	if p.TypeID1 == 0 {
		err = errors.New("一级分类ID不能为空")
		return
	}

	if p.TypeID2 == 0 {
		err = errors.New("二级分类ID不能为空")
		return
	}

	if p.SecurityTypeID == 0 {
		err = errors.New("安保形式ID不能为空")
		return
	}

	return
}

