package main

import (
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

// Gin 初始化
func init() {
	router = gin.Default()

}

// InitHandler 初始化配置
func InitHandler() *gin.Engine {
	return router
}
