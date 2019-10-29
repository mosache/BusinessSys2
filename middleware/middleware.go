package middleware

import "github.com/gin-gonic/gin"

type middeWare struct {
	HandleFunc gin.HandlerFunc
	Weight     int
}

type middleWareSlice []middeWare

var mws middleWareSlice

func (m middleWareSlice) Len() int { return len(m) }

func (m middleWareSlice) Less(i, j int) bool { return m[i].Weight > m[j].Weight }

func (m middleWareSlice) Swap(i, j int) { m[i], m[j] = m[j], m[i] }
