package dao

import (
	"iot_backend/param"
	"iot_backend/util"
)

// QueryHistoryData
// @Description: 查询电池历史状态
// @param data 绑定的 History 结构体数组指针
// @param param 前端传入的参数
// @return int64 查询到的数据条数
func QueryHistoryData(data interface{}, param param.HistoryParam) int64 {
	db, _ := util.GetOrm()
	result := db.Where("rfid = ?", param.ID).Where("created BETWEEN ? AND ?", param.StartTime, param.EndTime).Find(data)
	counts := result.RowsAffected

	//TODO 分页功能

	return counts
}
