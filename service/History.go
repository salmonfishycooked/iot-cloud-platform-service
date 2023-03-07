package service

import (
	"iot_backend/dao"
	"iot_backend/model"
	"iot_backend/param"
)

// GetHistorySensor
// @Description: 获取历史传感数据
// @param data 绑定的 Model 层数据
// @param par
// @return int64 数据条数
func GetHistorySensor(data *[]model.HistorySensorData, par param.HistorySensorParam) int64 {
	counts := dao.QueryHistorySensor(data, par.DeviceTag, par.SensorTag)

	return counts
}

// GetHistoryOrder
// @Description: 获取历史命令数据
// @param data 绑定的 Model 层数据
// @param par
// @return int64 数据条数
func GetHistoryOrder(data *[]model.HistoryOrderData, par param.HistoryActuatorParam) int64 {
	counts := dao.QueryHistoryActuator(data, par.DeviceTag, par.ActuatorTag)

	return counts
}

// CreateHistorySensor
// @Description: 创建历史传感数据
// @param data
// @return error
func CreateHistorySensor(data model.HistorySensorData) {
	// 检测该传感器是否存在
	counts := dao.QuerySensorSingle(&model.Sensor{}, data.DeviceTag, data.SensorTag)
	if counts == 0 {
		return
	}

	dao.CreateData(&data)
}

// CreateHistoryOrder
// @Description: 创建历史命令数据
// @param data
// @return error
func CreateHistoryOrder(data model.HistoryOrderData) {
	// 检测该执行器是否存在
	counts := dao.QueryActuatorSingle(&model.Actuator{}, data.DeviceTag, data.ActuatorTag)
	if counts == 0 {
		return
	}

	dao.CreateData(&data)
}
