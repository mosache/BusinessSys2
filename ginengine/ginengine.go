package ginengine

import (
	"sync"

	"github.com/gin-gonic/gin"

	config "BusinessSys/Config"
)

//Engine gin engine
var Engine *gin.Engine

var ginEngineOnec sync.Once

//InitEngine InitEngine
func InitEngine() {
	ginEngineOnec.Do(func() {
		Engine = gin.New()
		if config.AppConfig.Mode == "debug" {
			gin.SetMode(gin.DebugMode)
		} else {
			gin.SetMode(gin.ReleaseMode)
		}
	})
}
