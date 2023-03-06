package service

import (
	"iot_backend/dao"
	"iot_backend/model"
	"iot_backend/param"
)

// GetHistoryList
// @Description: 获取电池历史数据列表
// @param history
// @param param
// @return int64
func GetHistoryList(history []model.History, param param.HistoryParam) int64 {
	counts := dao.QueryHistoryData(&history, param)
	return counts
}
