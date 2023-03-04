package service

import (
	"iot_backend/dao"
	"iot_backend/model"
)

// GetBatteryInfo
// @Description: 通过电池id获取电池最新数据
// @param id 电池id
func GetBatteryInfo(id uint) model.Battery {
	battery := dao.QueryBatteryById(id) // 以 id 的方式查找电池
	return battery
}
