package dao

import (
	"MyGo_middleware/common/config"
	"MyGo_middleware/common/logger"
	"go_http/entity/dto_entity"
)

func CreateDataInput(pid string, request []byte) error {
	var csData dto_entity.DataInput
	csData.Pid = pid
	csData.DataRecord = string(request)
	err := config.DbCon.Table("data_input").Where("pid = ?", pid).Create(&csData).Error
	if err != nil {
		logger.BuilderWithTraceId(pid).Business().Field("", "").Build().ToErrorLog()
		return err
	}
	return nil
}
