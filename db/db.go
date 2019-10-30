package db

import (
	"BusinessSys/config"
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
)

var (
	DB *gorm.DB
)

func InitDB() {
	dbStr := fmt.Sprintf("%s:%s@%s/%s?charset=utf8", config.AppConfig.Username, config.AppConfig.Password, config.AppConfig.Addr, config.AppConfig.DataBaseName)

	var err error

	DB, err = gorm.Open("mysql", dbStr)

	if err != nil {
		log.Fatal(err.Error())
	}

}
