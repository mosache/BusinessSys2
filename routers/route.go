package routes

import (
	"log"
	"sync"

	"github.com/gin-gonic/gin"
)

type routeFunc struct {
	Name   string
	Weight int
	Func   func(e *gin.Engine)
}

type routeFuncSlice []routeFunc

func (r routeFuncSlice) Len() int { return len(r) }

func (r routeFuncSlice) Less(i, j int) bool { return r[i].Weight > r[j].Weight }

func (r routeFuncSlice) Swap(i, j int) { r[i], r[j] = r[j], r[i] }

var userRouteOnce sync.Once
var routes routeFuncSlice

//注册新route
//name 用来防止重复注册route
func register(name string, f func(e *gin.Engine)) {
	registerWithWeight(name, 50, f)
}

func registerWithWeight(name string, weight int, f func(e *gin.Engine)) {
	if weight < 0 || weight > 100 {
		log.Fatal("")
	}
}
