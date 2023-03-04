package controller

import (
	"github.com/gin-gonic/gin"
	"iot_backend/model"
	"iot_backend/param"
	"iot_backend/service"
	"iot_backend/util"
)

// getHistoryData
// @Description: 获取电池历史数据
// @param ctx 上下文
func getHistoryData(ctx *gin.Context) {
	param := param.HistoryParam{}
	ctx.ShouldBindJSON(&param)

	// 传入 ID 有误或者没传
	if param.ID == 0 {
		util.ResponseError(ctx)
		return
	}

	var data []model.History
	counts := service.GetHistoryList(&data, param) // 获取信息

	// 如果没有查询到数据
	if counts == 0 {
		util.ResponseErrorWithMsg(ctx, "未查询到相关数据")
		return
	}
	// 正常返回
	util.ResponseOK(ctx, data)
}

// InitHistoryRoutes
// @Description: 初始化 History 的路由
// @param group 接收 app 的路由分组
func InitHistoryRoutes(group *gin.RouterGroup) {
	group.POST("/history", getHistoryData)
}
