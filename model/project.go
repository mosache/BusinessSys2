package model

import "time"

type Project struct {
	ProjectID   int64  `gorm:"column:id;primary_key" `
	ProjectName string `gorm:"column:project_name"`
	Address string `gorm:"column:address"`
	ProvinceName string `gorm:"column:province_name"`
	ProvinceID int64 `gorm:"column:province_id"`
	CityName string `gorm:"column:city_name"`
	CityID int64 `gorm:"column:city_id"`
	AreaName string `gorm:"column:area_name"`
	AreaID int64 `gorm:"column:area_id"`
	CoordinateX string `gorm:"column:coordinate_x"`
	CoordinateY string `gorm:"column:coordinate_y"`
	TypeID1 int64 `gorm:"column:type_id_1"`
	TypeID2 int64 `gorm:"column:type_id_2"`
	SecurityTypeID int64 `gorm:"column:security_type_id"`
	AddUserID int64 `gorm:"column:add_user_id"` //'创建人',
	AddTime time.Time `gorm:"column:add_time""` //'添加时间'
}

func (Project) TableName() string {
	return "project_base_info"
}