package business

import (
	"encoding/json"
	"go_http/dao"
	"go_http/entity/api_entity"
	"go_http/entity/engine_entity"
)

func ProductInfo(request api_entity.ProductData) error {
	//logger.BuilderWithNotCtx().Business().Field("ProductInfo", "start").Build().ToInfoLog()
	var result engine_entity.Result
	result.ProductName = request.ReqData.ProductName
	result.ProductType = request.ReqData.ProductType
	result.ProductNum = request.ReqData.ProductNum

	byteData, err := json.Marshal(result)
	if err != nil {
		//logger.BuilderWithTraceId(request.Pid).Business().Field("Marshal", "failed").Build().ToErrorLog()
		return err
	}

	err = dao.CreateDataInput(request.Pid, byteData)
	if err != nil {
		//logger.BuilderWithTraceId(request.Pid).Business().Field("CreateDataInput", "failed").Build().ToErrorLog()
		return err
	}

	//logger.BuilderWithTraceId(request.Pid).Business().Field("ProductInfo", "finish").Build().ToErrorLog()
	return nil
}
