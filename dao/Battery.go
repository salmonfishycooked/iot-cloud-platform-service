package dao

import (
	"iot_backend/model"
	"iot_backend/util"
)

// QueryBatteryById
// @Description: 以 id 的方式查找电池
// @param id 电池id
// @return model.Battery 查找到的电池结构体
func QueryBatteryById(id uint) model.Battery {
	db, _ := util.GetOrm()
	battery := model.Battery{}
	db.First(&battery, id)

	return battery
}
