package controller

import (
	"github.com/gin-gonic/gin"
	"iot_backend/model"
	"iot_backend/param"
	"iot_backend/service"
	"iot_backend/util"
)

// InitSensorRoutes
// @Description: 初始化 Sensor 的路由
// @param group 接收 app 的路由分组
func InitSensorRoutes(group *gin.RouterGroup) {
	device := group.Group("/sensor")
	{
		device.POST("/info", QuerySensor)
		device.POST("/update", updateSensor)
		device.POST("/create", createSensor)
		device.POST("/delete", deleteSensor)
	}
}

func QuerySensor(ctx *gin.Context) {
	queryParam := param.SensorQueryParam{}
	ctx.ShouldBindJSON(&queryParam)

	if queryParam.Tag == "" || queryParam.DeviceTag == "" {
		util.ResponseErrorWithMsg(ctx, "输入数据有误！")
		return
	}

	data := model.Sensor{}
	counts := service.QuerySensor(&data, queryParam)

	// 如果没有查询到数据
	if counts == 0 {
		util.ResponseErrorWithMsg(ctx, "未查询到相关数据")
		return
	}
	// 正常返回
	util.ResponseOK(ctx, data)
}

// createSensor
// @Description: 新建传感器
// @param ctx
func createSensor(ctx *gin.Context) {
	createParam := param.SensorCreateParam{}
	ctx.ShouldBindJSON(&createParam)

	if createParam.Tag == "" || createParam.Name == "" || createParam.DeviceTag == "" {
		util.ResponseErrorWithMsg(ctx, "输入数据有误！")
		return
	}

	err := service.CreateSensor(createParam) // 创建数据
	if err != nil {
		util.ResponseErrorWithMsg(ctx, "输入数据有误或设备tag不存在或者传感器tag已被占用！")
		return
	}
	// 正常返回
	util.ResponseOK(ctx, nil)
}

func updateSensor(ctx *gin.Context) {
	updateParam := param.SensorUpdateParam{}
	ctx.ShouldBindJSON(&updateParam)

	if updateParam.Tag == "" || updateParam.DeviceTag == "" {
		util.ResponseErrorWithMsg(ctx, "输入数据有误！")
		return
	}

	err := service.UpdateSensor(updateParam)
	if err != nil {
		util.ResponseErrorWithMsg(ctx, "不存在该传感器！")
		return
	}
	// 正常返回
	util.ResponseOK(ctx, nil)
}

func deleteSensor(ctx *gin.Context) {
	queryParam := param.SensorQueryParam{}
	ctx.ShouldBindJSON(&queryParam)

	if queryParam.Tag == "" || queryParam.DeviceTag == "" {
		util.ResponseErrorWithMsg(ctx, "输入数据有误！")
		return
	}

	err := service.DeleteSensor(queryParam)
	if err != nil {
		util.ResponseErrorWithMsg(ctx, "不存在该传感器！")
		return
	}
	// 正常返回
	util.ResponseOK(ctx, nil)
}
