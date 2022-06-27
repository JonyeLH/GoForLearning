package controller

import (
	"MyGo_middleware/common/logger"
	"MyGo_middleware/common/panic_handler"
	"MyGo_middleware/constant"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go_http/business"
	"go_http/entity/api_entity"
	"net/http"
)

type Base struct {
}

var HttpBase Base

func (*Base) LogTest(ctx *gin.Context) {

}

func (*Base) MysqlTest(ctx *gin.Context) {
	defer panic_handler.PanicHandler(nil)
	var request api_entity.ProductData
	var result api_entity.BaseHttpInfo

	if err := ctx.ShouldBindBodyWith(&request, binding.JSON); err != nil {
		result.Code = constant.Failed
		result.Message = "parameter parse failed"
		logger.BuilderWithNotCtx().Business().Field("ShouldBindBodyWith", "failed").Build().ToErrorLog()
		ctx.JSON(http.StatusOK, result)
		return
	}

	err := business.ProductInfo(request)
	if err != nil {
		logger.BuilderWithTraceId(request.Pid).Business().Field("ProductInfo", "failed").Build().ToInfoLog()
	}

	result.Code = constant.Success
	result.Message = "success"
	logger.BuilderWithTraceId(request.Pid).Business().Field("MysqlTest", "success").Build().ToInfoLog()
	ctx.JSON(http.StatusOK, result)
	return

}
