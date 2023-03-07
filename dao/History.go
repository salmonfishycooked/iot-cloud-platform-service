package dao

import (
	"iot_backend/model"
	"iot_backend/util"
)

// QueryHistorySensor
// @Description: 获取历史传感数据
// @param data 绑定的 Model 层数据
// @param deviceTag 设备tag
// @param sensorTag 传感器tag
// @return int64 数据条数
func QueryHistorySensor(data *[]model.HistorySensorData, deviceTag string, sensorTag string) int64 {
	db, _ := util.GetOrm()
	result := db.Where("device_tag = ? AND sensor_tag = ?", deviceTag, sensorTag).Find(&data)
	counts := result.RowsAffected

	return counts
}

// QueryHistoryActuator
// @Description: 获取历史命令数据
// @param data 绑定的 Model 层数据
// @param deviceTag 设备tag
// @param ActuatorTag 执行器tag
// @return int64 数据条数
func QueryHistoryActuator(data *[]model.HistoryOrderData, deviceTag string, ActuatorTag string) int64 {
	db, _ := util.GetOrm()
	result := db.Where("device_tag = ? AND actuator_tag = ?", deviceTag, ActuatorTag).Find(&data)
	counts := result.RowsAffected

	return counts
}
