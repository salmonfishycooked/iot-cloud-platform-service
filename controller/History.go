package controller

import (
	"github.com/gin-gonic/gin"
	"iot_backend/model"
	"iot_backend/param"
	"iot_backend/service"
	"iot_backend/util"
)

// InitHistoryRoutes
// @Description: 初始化 History 的路由
// @param group 接收 app 的路由分组
func InitHistoryRoutes(group *gin.RouterGroup) {
	device := group.Group("/history")
	{
		device.POST("/sensor", getHistorySensor)
		device.POST("/order", getHistoryOrder)
	}
}

// getHistorySensor
// @Description: 获取历史传感数据
// @param ctx
func getHistorySensor(ctx *gin.Context) {
	historyParam := param.HistorySensorParam{}
	ctx.ShouldBindJSON(&historyParam)

	if historyParam.SensorTag == "" || historyParam.DeviceTag == "" || historyParam.StartTime == 0 || historyParam.EndTime == 0 {
		util.ResponseErrorWithMsg(ctx, "输入数据有误！")
		return
	}

	var data []model.HistorySensorData
	counts := service.GetHistorySensor(&data, historyParam)
	if counts == 0 {
		util.ResponseErrorWithMsg(ctx, "未查询到相关数据！")
		return
	}
	// 正常返回
	util.ResponseOK(ctx, data)
}

// getHistoryOrder
// @Description: 获取历史命令数据
// @param ctx
func getHistoryOrder(ctx *gin.Context) {
	historyParam := param.HistoryActuatorParam{}
	ctx.ShouldBindJSON(&historyParam)

	if historyParam.ActuatorTag == "" || historyParam.DeviceTag == "" || historyParam.StartTime == 0 || historyParam.EndTime == 0 {
		util.ResponseErrorWithMsg(ctx, "输入数据有误！")
		return
	}

	var data []model.HistoryOrderData
	counts := service.GetHistoryOrder(&data, historyParam)
	if counts == 0 {
		util.ResponseErrorWithMsg(ctx, "未查询到相关数据！")
		return
	}
	// 正常返回
	util.ResponseOK(ctx, data)
}
