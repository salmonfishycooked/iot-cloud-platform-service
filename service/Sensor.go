package service

import (
	"github.com/gin-gonic/gin"
	"iot_backend/dao"
	"iot_backend/model"
	"iot_backend/param"
)

// CreateSensor
// @Description: 创建数据
// @param par
// @return error
func CreateSensor(par param.SensorCreateParam) error {
	data := model.Sensor{
		Tag:       par.Tag,
		Name:      par.Name,
		DeviceTag: par.DeviceTag,
	}

	// 检测 DeviceTag 在设备表里是否存在
	counts := dao.QueryByTag(&model.Device{}, par.DeviceTag)
	if counts == 0 {
		return gin.Error{}
	}

	err := dao.CreateData(&data) // 创建数据
	return err
}

// UpdateSensor
// @Description: 更新传感器数据
// @param par
// @return error
func UpdateSensor(par param.SensorUpdateParam) error {
	err, count := dao.UpdateSensor(par.DeviceTag, par.Tag, par.Value)

	// 如果没有找到该传感器
	if count == 0 {
		return gin.Error{}
	}

	return err
}
