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

	// 在历史传感数据里面新增一条记录
	historySensor := model.HistorySensorData{
		DeviceTag: par.DeviceTag,
		SensorTag: par.Tag,
		Value:     par.Value,
	}
	CreateHistorySensor(historySensor)

	return err
}

// QuerySensor
// @Description: 查询某个传感器的数据
// @param data 绑定的 Model 层数据
// @param par 查询参数
// @return int64 数据条数
func QuerySensor(data *model.Sensor, par param.SensorQueryParam) int64 {
	count := dao.QuerySensorSingle(data, par.DeviceTag, par.Tag)

	return count
}

// DeleteSensor
// @Description: 删除传感器
// @param par 查询参数
// @return error
func DeleteSensor(par param.SensorQueryParam) error {
	err, counts := dao.DeleteSensor(par.DeviceTag, par.Tag)

	// 未找到传感器
	if counts == 0 {
		return gin.Error{}
	}

	return err
}
