package model

import "time"

type User struct {
	UserID     int64     `gorm:"column:user_id;primary_key"`
	UserName   string    `gorm:"column:username"`
	Password   string    `gorm:"column:password"`
	CreateTime time.Time `gorm:"column:create_time"`
}
