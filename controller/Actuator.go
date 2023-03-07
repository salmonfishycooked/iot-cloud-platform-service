package controller

import (
	"github.com/gin-gonic/gin"
	"iot_backend/model"
	"iot_backend/param"
	"iot_backend/service"
	"iot_backend/util"
)

// InitActuatorRoutes
// @Description: 初始化 Actuator 的路由
// @param group 接收 app 的路由分组
func InitActuatorRoutes(group *gin.RouterGroup) {
	device := group.Group("/actuator")
	{
		device.POST("/order", orderActuator)
		device.POST("/create", createActuator)
		device.POST("/delete", deleteActuator)
		device.POST("/info", QueryActuator)
	}
}

// createActuator
// @Description: 新建执行器
// @param ctx
func createActuator(ctx *gin.Context) {
	createParam := param.ActuatorCreateParam{}
	ctx.ShouldBindJSON(&createParam)

	if createParam.Tag == "" || createParam.Name == "" || createParam.DeviceTag == "" {
		util.ResponseErrorWithMsg(ctx, "输入数据有误！")
		return
	}

	err := service.CreateActuator(createParam) // 创建数据
	if err != nil {
		util.ResponseErrorWithMsg(ctx, "输入数据有误或设备tag不存在或者执行器tag已被占用！")
		return
	}
	// 正常返回
	util.ResponseOK(ctx, nil)
}

// updateActuator
// @Description: 执行器下发命令
// @param ctx
func orderActuator(ctx *gin.Context) {
	orderParam := param.ActuatorOrderParam{}
	ctx.ShouldBindJSON(&orderParam)
	if orderParam.Value == "" || orderParam.DeviceTag == "" || orderParam.Tag == "" {
		util.ResponseErrorWithMsg(ctx, "输入数据有误！")
		return
	}

	err := service.OrderActuator(orderParam) // 发命令
	if err != nil {
		util.ResponseErrorWithMsg(ctx, "发送失败！请检查该执行器是否存在或者设备是否上线")
		return
	}
	// 正常返回
	util.ResponseOKWithMsg(ctx, nil, "命令已发出！")
}

// QueryActuator
// @Description: 获取执行器信息
// @param ctx
func QueryActuator(ctx *gin.Context) {
	queryParam := param.ActuatorQueryParam{}
	ctx.ShouldBindJSON(&queryParam)

	if queryParam.Tag == "" || queryParam.DeviceTag == "" {
		util.ResponseErrorWithMsg(ctx, "输入数据有误！")
		return
	}

	data := model.Actuator{}
	counts := service.QueryActuator(&data, queryParam)

	// 如果没有查询到数据
	if counts == 0 {
		util.ResponseErrorWithMsg(ctx, "未查询到相关数据")
		return
	}
	// 正常返回
	util.ResponseOK(ctx, data)
}

// deleteActuator
// @Description: 删除传感器
// @param ctx
func deleteActuator(ctx *gin.Context) {
	queryParam := param.ActuatorQueryParam{}
	ctx.ShouldBindJSON(&queryParam)

	if queryParam.Tag == "" || queryParam.DeviceTag == "" {
		util.ResponseErrorWithMsg(ctx, "输入数据有误！")
		return
	}

	err := service.DeleteActuator(queryParam)
	if err != nil {
		util.ResponseErrorWithMsg(ctx, "不存在该执行器！")
		return
	}
	// 正常返回
	util.ResponseOK(ctx, nil)
}
