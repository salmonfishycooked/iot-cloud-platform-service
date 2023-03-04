package dao

import (
	"iot_backend/model"
	"iot_backend/util"
)

// QueryDeviceById
// @Description: 以设备id来查询设备最新数据
// @param id 设备 id
// @return model.Battery 查询到的数据结构体
func QueryDeviceById(id uint) model.Device {
	db, _ := util.GetOrm()
	data := model.Device{}
	db.First(&data, id)

	return data
}
