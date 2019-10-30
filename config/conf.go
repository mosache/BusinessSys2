package config

import (
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

//Config Config
type Config struct {
	Server

	DataBase
}

type Server struct {
	Port int
	Mode string
}

type DataBase struct {
	Addr string

	Username string

	Password string

	DataBaseName string
}

//AppConfig 配置项
var AppConfig *Config

//InitConfig 初始化配置
func InitConfig() {
	configFile, err := os.OpenFile("./appConfig.yaml", os.O_RDONLY, 0222)

	if err != nil {
		//配置文件读取出错或不存在,使用默认配置
		AppConfig = getDefault()
		return
	}

	defer configFile.Close()

	var bytes []byte

	bytes, err = ioutil.ReadAll(configFile)

	if err != nil {
		//配置文件读取出错或不存在,使用默认配置
		AppConfig = getDefault()
		return
	}

	err = yaml.Unmarshal(bytes, AppConfig)

	if err != nil {
		//配置文件读取出错或不存在,使用默认配置
		AppConfig = getDefault()
		return
	}
}

func getDefault() *Config {
	return &Config{
		Server{Port: 8081,
			Mode: "debug"},
		DataBase{
			Addr:         "192.168.66.101:3307",
			Username:     "root",
			Password:     "123456",
			DataBaseName: "",
		},
	}
}
