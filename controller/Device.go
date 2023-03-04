package controller

import (
	"github.com/gin-gonic/gin"
	"iot_backend/service"
	"iot_backend/util"
)

// DeviceParam
// @Description: 前端发来的参数结构体模型
type DeviceParam struct {
	ID uint `json:"id"`
}

// getDeviceData
// @Description: 获取设备最新状态
// @param ctx 上下文
func getDeviceData(ctx *gin.Context) {
	deviceParam := DeviceParam{}
	ctx.ShouldBindJSON(&deviceParam)

	// 传入 ID 有误或者没传
	if deviceParam.ID == 0 {
		util.ResponseError(ctx)
		return
	}

	data := service.GetDeviceInfo(deviceParam.ID) // 获取设备信息
	util.ResponseOK(ctx, data)
}

// InitDeviceRoutes
// @Description: 初始化 Device 的路由
// @param group 接收 app 的路由分组
func InitDeviceRoutes(group *gin.RouterGroup) {
	group.POST("/device", getDeviceData)
}
