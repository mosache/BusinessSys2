package ginengine

import (
	"sync"

	"github.com/gin-gonic/gin"

	"BusinessSys/config"
)

//Engine gin engine
var Engine *gin.Engine

var ginEngineOnce sync.Once

//InitEngine InitEngine
func InitEngine() {
	ginEngineOnce.Do(func() {
		if config.AppConfig.Mode == "debug" {
			gin.SetMode(gin.DebugMode)
		} else {
			gin.SetMode(gin.ReleaseMode)
		}
		Engine = gin.New()
	})
}
