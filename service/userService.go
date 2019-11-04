package service

import "BusinessSys/db"

func LoginIn(username string, password string) (*db.User, bool) {
	user := &db.User{}

	db.DB.First(user, "username = ?", username)

	if user.Password == password {
		return user, true
	}

	return nil, false
}
