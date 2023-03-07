package dao

import (
	"iot_backend/model"
	"iot_backend/util"
)

// UpdateDeviceStatus
// @Description: 更新设备在线状态
// @param tag 设备tag
// @param isOnline 在线状态
// @return error
// @return int64 数据条数
func UpdateDeviceStatus(tag string, isOnline bool) (error, int64) {
	db, _ := util.GetOrm()
	var count int64
	result := db.Model(&model.Device{}).Where("tag = ?", tag).Count(&count).Update("is_online", isOnline)
	err := result.Error

	return err, count
}
