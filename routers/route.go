package routes

import (
	"BusinessSys/ginengine"
	"log"
	"net/http"
	"sort"
	"strings"
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
//默认route weight 为 50
func register(name string, f func(e *gin.Engine)) {
	registerWithWeight(name, 50, f)
}

func registerWithWeight(name string, weight int, f func(e *gin.Engine)) {
	if weight < 0 || weight > 100 {
		log.Fatalf("weight should be 0 <= weight <= 100 with route : %s", name)
	}

	for _, r := range routes {
		if strings.ToLower(name) == strings.ToLower(r.Name) {
			log.Fatalf("has two routes has same name : %s", name)
		}
	}

	routes = append(routes, routeFunc{Name: name, Weight: weight, Func: f})
}

//Setup framework init
func Setup() {
	sort.Sort(routes)

	for _, r := range routes {
		r.Func(ginengine.Engine)
		log.Printf("add route of %s", r.Name)
	}
}

func ResponseOk(data interface{}) map[string]interface{} {
	return gin.H{
		"status": http.StatusOK,
		"msg":    "",
		"data":   data,
	}
}

func ResponseBad(msg string) map[string]interface{} {
	return gin.H{
		"status": http.StatusBadRequest,
		"msg":    msg,
	}
}
