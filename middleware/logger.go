package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
)

func init() {
	registerWithWeight(100, Logger)
}

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Printf("[%s] %s \n", c.Request.Method, c.Request.URL.String())
		c.Next()
	}
}
