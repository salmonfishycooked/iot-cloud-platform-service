package controller

import (
	"github.com/gin-gonic/gin"
	"iot_backend/param"
	"iot_backend/service"
	"iot_backend/util"
)

// InitDeviceRoutes
// @Description: 初始化 Device 的路由
// @param group 接收 app 的路由分组
func InitDeviceRoutes(group *gin.RouterGroup) {
	device := group.Group("/device")
	{
		device.POST("/info", getDeviceInfo)
		device.POST("/create", createDevice)
	}
}

// getDeviceData
// @Description: 获取设备最新状态
// @param ctx 上下文
func getDeviceInfo(ctx *gin.Context) {
	infoParam := param.DeviceParam{}
	ctx.ShouldBindJSON(&infoParam)

	// 传入 ID 有误或者没传
	if infoParam.DeviceTag == "" {
		util.ResponseError(ctx)
		return
	}

	data := param.DeviceResponse{}
	counts := service.QueryDeviceByTag(&data, infoParam.DeviceTag) // 获取设备信息

	// 如果没有查询到数据
	if counts == 0 {
		util.ResponseErrorWithMsg(ctx, "未查询到相关数据")
		return
	}
	// 正常返回
	util.ResponseOK(ctx, data)
}

// createDevice
// @Description: 新建设备
// @param ctx
func createDevice(ctx *gin.Context) {
	createParam := param.DeviceCreateParam{}
	ctx.ShouldBindJSON(&createParam)

	if createParam.Tag == "" || createParam.Name == "" {
		util.ResponseErrorWithMsg(ctx, "输入数据有误！")
		return
	}

	err := service.CreateDevice(createParam) // 创建数据
	if err != nil {
		util.ResponseErrorWithMsg(ctx, "输入数据有误或设备tag已存在！")
		return
	}
	// 正常返回
	util.ResponseOK(ctx, nil)
}
