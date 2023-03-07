package service

import (
	"github.com/gin-gonic/gin"
	"iot_backend/dao"
	"iot_backend/model"
	"iot_backend/param"
)

// QueryDeviceByTag
// @Description: 返回设备以及它所有的传感器数据（以 Tag 查询）
// @param res
// @param tag 传入的 Tag
// @return int64 查询到的数据条数
func QueryDeviceByTag(res *param.DeviceResponse, tag string) int64 {
	device := model.Device{}
	var sensors []model.Sensor

	counts := dao.QueryByTag(&device, tag)    // 查询设备
	dao.QuerySensorByDeviceTag(&sensors, tag) // 查询该设备的所有传感器

	res.Device = device
	res.Sensors = sensors

	return counts
}

// CreateDevice
// @Description: 创建数据
// @param par
// @return error
func CreateDevice(par param.DeviceCreateParam) error {
	data := model.Device{
		Tag:  par.Tag,
		Name: par.Name,
	}
	err := dao.CreateData(&data) // 创建数据
	return err
}

// DeleteDevice
// @Description: 删除设备
// @param par
// @return error
func DeleteDevice(par param.DeviceParam) error {
	var sensors []model.Sensor
	counts := dao.QuerySensorByDeviceTag(&sensors, par.Tag) // 查询该设备的所有传感器
	// 该设备下还存在有传感器，不能删除
	if counts != 0 {
		return gin.Error{}
	}

	// 不存在该设备
	err, counts := dao.DeleteByTag(&model.Device{}, par.Tag)
	if counts == 0 {
		return gin.Error{}
	}
	return err
}
