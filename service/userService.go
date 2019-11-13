package service

import (
	"BusinessSys/db"
	"BusinessSys/model"
	"strings"
)

func LoginIn(username string, password string) (*model.User, bool) {
	user := &model.User{}

	db.DB.First(user, "username = ?", username)

	if strings.ToLower(user.Password)  == strings.ToLower(password) {
		return user, true
	}

	return nil, false
}

/**
通过ID获取User
 */
func GetUserByID(userID int64) (user *model.User,err error){

	user = &model.User{}

	if err = db.DB.First(&user).Where("user_id = ?",userID).Error;err!=nil {
		return
	}

	return
}

