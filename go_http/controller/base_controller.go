package controller

import "github.com/gin-gonic/gin"

type Base struct {
}

var HttpBase Base

func (*Base) LogTest(ctx *gin.Context) {

}
