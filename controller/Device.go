package controller

import (
	"github.com/gin-gonic/gin"
	"iot_backend/model"
	"iot_backend/param"
	"iot_backend/service"
	"iot_backend/util"
)

// getDeviceData
// @Description: 获取设备最新状态
// @param ctx 上下文
func getDeviceData(ctx *gin.Context) {
	deviceParam := param.DeviceParam{}
	ctx.ShouldBindJSON(&deviceParam)

	// 传入 ID 有误或者没传
	if deviceParam.ID == 0 {
		util.ResponseError(ctx)
		return
	}

	data := model.Device{}
	counts := service.GetInfoById(&data, deviceParam.ID) // 获取电池信息

	// 如果没有查询到数据
	if counts == 0 {
		util.ResponseErrorWithMsg(ctx, "未查询到相关数据")
		return
	}
	// 正常返回
	util.ResponseOK(ctx, data)
}

// InitDeviceRoutes
// @Description: 初始化 Device 的路由
// @param group 接收 app 的路由分组
func InitDeviceRoutes(group *gin.RouterGroup) {
	group.POST("/device", getDeviceData)
}
