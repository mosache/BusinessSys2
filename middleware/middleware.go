package middleware

import (
	"BusinessSys/ginengine"
	"github.com/gin-gonic/gin"
	"log"
	"sort"
	"sync"
)

type middleWare struct {
	HandleFunc func() gin.HandlerFunc
	Weight     int
}

type middleWareSlice []middleWare

func (m middleWareSlice) Len() int { return len(m) }

func (m middleWareSlice) Less(i, j int) bool { return m[i].Weight > m[j].Weight }

func (m middleWareSlice) Swap(i, j int) { m[i], m[j] = m[j], m[i] }

var mws middleWareSlice

var middleWareOnce sync.Once

//注册中间件
//
func register(f func() gin.HandlerFunc) {
	registerWithWeight(50, f)
}

func registerWithWeight(weight int, f func() gin.HandlerFunc) {
	if weight < 0 || weight > 100 {
		log.Fatal("middleware weight should be 0 <= weight <= 100 with")
	}

	mw := middleWare{
		HandleFunc: f,
		Weight:     weight,
	}

	mws = append(mws, mw)
}

//SetUp init middleWares
func SetUp() {
	sort.Sort(mws)

	for _, mw := range mws {
		ginengine.Engine.Use(mw.HandleFunc())
	}

	log.Println("load middleWareSuccess")

}
