package dao

import (
	"iot_backend/model"
	"iot_backend/util"
)

// QueryActuatorByDeviceTag
// @Description: 以设备tag查询该设备所拥有的所有执行器信息
// @param data 需要绑定的数据
// @param tag 传入设备tag
// @return int64 查询到的数据条数
func QueryActuatorByDeviceTag(data *[]model.Actuator, tag string) int64 {
	db, _ := util.GetOrm()
	result := db.Where("device_tag = ?", tag).Find(&data)
	counts := result.RowsAffected

	return counts
}

// DeleteActuator
// @Description: 删除执行器
// @param deviceTag 该执行器所属设备的tag
// @param tag 执行器tag
// @return int64 数据条数
// @return error
func DeleteActuator(deviceTag string, tag string) (error, int64) {
	db, _ := util.GetOrm()
	result := db.Where("device_tag = ? AND tag = ?", deviceTag, tag).Delete(&model.Actuator{})
	err := result.Error
	count := result.RowsAffected

	return err, count
}
