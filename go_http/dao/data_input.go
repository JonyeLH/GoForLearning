package dao

import (
	"MyGo_middleware/common/config"
	"go_http/entity/dto_entity"
)

func CreateDataInput(pid string, request []byte) error {
	var csData dto_entity.DataInput
	csData.Pid = pid
	csData.DataRecord = string(request)

	err := config.DbCon.Table("data_input").Create(csData).Error
	if err != nil {
		//logger.BuilderWithTraceId(pid).Business().Field("", "").Build().ToErrorLog()
		return err
	}
	return nil
}

func UpdateDataBase(pid string, csData string) error {
	err := config.DbCon.Table("data_input").Where("pid =?", pid).Update("data_record", csData).Error
	if err != nil {
		//logger.BuilderWithTraceId(pid).Business().Field("", "").Build().ToErrorLog()
		return err
	}
	return nil
}
