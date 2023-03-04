package controller

import (
	"github.com/gin-gonic/gin"
	"iot_backend/model"
	"iot_backend/param"
	"iot_backend/service"
	"iot_backend/util"
)

// getBatteryData
// @Description: 获取电池最新状态
// @param ctx 上下文
func getBatteryData(ctx *gin.Context) {
	batteryParam := param.BatteryParam{}
	ctx.ShouldBindJSON(&batteryParam)

	// 传入 ID 有误或者没传
	if batteryParam.ID == 0 {
		util.ResponseError(ctx)
		return
	}

	data := model.Battery{}
	counts := service.GetInfoById(&data, batteryParam.ID) // 获取电池信息

	// 如果没有查询到数据
	if counts == 0 {
		util.ResponseErrorWithMsg(ctx, "未查询到相关数据")
		return
	}
	// 正常返回
	util.ResponseOK(ctx, data)
}

// InitBatteryRoutes
// @Description: 初始化 Battery 的路由
// @param group 接收 app 的路由分组
func InitBatteryRoutes(group *gin.RouterGroup) {
	group.POST("/battery", getBatteryData)
}
