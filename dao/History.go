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
// @param page 分页
// @param pageSize 分页
// @return int64 数据条数
func QueryHistorySensor(data *[]model.HistorySensorData, deviceTag string, sensorTag string, startTime int64, endTime int64) int64 {
	db, _ := util.GetOrm()
	result := db.Where("device_tag = ? AND sensor_tag = ?", deviceTag, sensorTag).Where("created >= ? AND created <= ?", startTime, endTime).Order("created desc").Find(&data)
	counts := result.RowsAffected

	return counts
}

// DeleteHistorySensor
// @Description: 删除历史传感数据
// @param sensorTag 对应的传感器tag
// @return int64 删除的数据条数
func DeleteHistorySensor(sensorTag string) int64 {
	db, _ := util.GetOrm()
	result := db.Where("sensor_tag = ?", sensorTag).Delete(&model.HistorySensorData{})
	counts := result.RowsAffected

	return counts
}

// QueryHistoryActuator
// @Description: 获取历史命令数据
// @param data 绑定的 Model 层数据
// @param deviceTag 设备tag
// @param ActuatorTag 执行器tag
// @return int64 数据条数
func QueryHistoryActuator(data *[]model.HistoryOrderData, deviceTag string, ActuatorTag string, startTime int64, endTime int64) int64 {
	db, _ := util.GetOrm()
	result := db.Where("device_tag = ? AND actuator_tag = ?", deviceTag, ActuatorTag).Where("created >= ? AND created <= ?", startTime, endTime).Find(&data)
	counts := result.RowsAffected

	return counts
}

// DeleteHistoryOrder
// @Description: 删除历史命令数据
// @param actuatorTag 对应的执行器tag
// @return int64 删除的数据条数
func DeleteHistoryOrder(actuatorTag string) int64 {
	db, _ := util.GetOrm()
	result := db.Where("actuator_tag = ?", actuatorTag).Delete(&model.HistoryOrderData{})
	counts := result.RowsAffected

	return counts
}
