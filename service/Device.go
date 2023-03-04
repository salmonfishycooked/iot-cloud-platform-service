package service

import (
	"iot_backend/dao"
	"iot_backend/model"
)

// GetDeviceInfo
// @Description: 通过设备id获取设备的最新状态
// @param id 设备id
// @return model.Device 查询到的设备数据结构体
func GetDeviceInfo(id uint) model.Device {
	data := dao.QueryDeviceById(id)
	return data
}
