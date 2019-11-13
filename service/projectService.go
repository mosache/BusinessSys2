package service

import (
	"BusinessSys/db"
	"BusinessSys/model"
)

/**
添加项目
 */
func InsertProject(project *model.Project) (err error) {

	if err = db.DB.Create(project).Error;err!=nil {
		return
	}
	return
}
