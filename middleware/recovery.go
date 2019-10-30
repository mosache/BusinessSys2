package middleware

import "github.com/gin-gonic/gin"

func init() {
	//recovery set weight max for it to use last
	registerWithWeight(100, func() gin.HandlerFunc {
		return gin.Recovery()
	})
}
