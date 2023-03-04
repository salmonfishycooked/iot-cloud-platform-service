package controller

import (
	"github.com/gin-gonic/gin"
	"iot_backend/service"
	"iot_backend/util"
)

// BatteryParam
// @Description: 前端发来的参数结构体模型
type BatteryParam struct {
	ID uint `json:"id"`
}

// getBatteryData
// @Description: 获取电池最新状态
// @param ctx 上下文
func getBatteryData(ctx *gin.Context) {
	batteryParam := BatteryParam{}
	ctx.ShouldBindJSON(&batteryParam)

	// 传入 ID 有误或者没传
	if batteryParam.ID == 0 {
		util.ResponseError(ctx)
		return
	}

	battery := service.GetBatteryInfo(batteryParam.ID) // 获取电池信息
	util.ResponseOK(ctx, battery)
}

// InitBatteryRoutes
// @Description: 初始化 Battery 的路由
// @param group 接收 app 的路由分组
func InitBatteryRoutes(group *gin.RouterGroup) {
	group.POST("/battery", getBatteryData)
}
