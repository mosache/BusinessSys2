package routes

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	_ "BusinessSys/docs"
)

func init() {
	register("swagger", func(e *gin.Engine) {
		e.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	})
}
