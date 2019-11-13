package model

import (
	"gopkg.in/go-playground/validator.v8"
)

type NoticePoint struct {
	ID int `gorm:"column:id"`
	Name string `gorm:"column:name" binding:"required"`
	TypeID int `gorm:"column:type_id" binding:"required,notZero"`
	CoordinateX string `gorm:"column:coordinate_x" binding:"required"`
	CoordinateY string `gorm:"column:coordinate_y" binding:"required"`
	FeelNow int `gorm:"column:feel_now" binding:"required"`
}

func (n *NoticePoint) GetError(err validator.ValidationErrors) string {

	if val,ok := err["NoticePoint.Name"];ok{
		if val.Field == "Name" {
			switch val.Tag {
			case "required":
				return "名称不能为空"
			}
		}
	}

	if val,ok := err["NoticePoint.TypeID"];ok{
		if val.Field == "TypeID" {
			switch val.Tag {
			case "required":
				return "类型不能为空"
			}
		}
	}

	if val,ok := err["NoticePoint.CoordinateX"];ok{
		if val.Field == "CoordinateX" {
			switch val.Tag {
			case "required":
				return "纬度不能为空"
			}
		}
	}

	if val,ok := err["NoticePoint.CoordinateY"];ok{
		if val.Field == "CoordinateY" {
			switch val.Tag {
			case "required":
				return "经度不能为空"
			}
		}
	}

	if val,ok := err["NoticePoint.FeelNow"];ok{
		if val.Field == "FeelNow" {
			switch val.Tag {
			case "required":
				return "当前印象不能为空"
			}
		}
	}

	return "参数错误"
}



