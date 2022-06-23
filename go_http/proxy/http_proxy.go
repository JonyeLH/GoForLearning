package proxy

import "github.com/gin-gonic/gin"

func NewGinRouter() *gin.Engine {
	ginRouter := gin.Default()
	return ginRouter
}
