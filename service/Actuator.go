package service

import (
	"github.com/gin-gonic/gin"
	"iot_backend/dao"
	"iot_backend/model"
	"iot_backend/param"
)

// CreateActuator
// @Description: 创建数据
// @param par
// @return error
func CreateActuator(par param.ActuatorCreateParam) error {
	data := model.Actuator{
		Tag:       par.Tag,
		Name:      par.Name,
		DeviceTag: par.DeviceTag,
	}

	// 检测 DeviceTag 在设备表里是否存在
	counts := dao.QueryByTag(&model.Device{}, par.DeviceTag)
	if counts == 0 {
		return gin.Error{}
	}

	err := dao.CreateData(&data) // 创建数据
	return err
}

// QueryActuator
// @Description: 查询某个执行器的信息
// @param data 绑定的 Model 层数据
// @param par 查询参数
// @return int64 数据条数
func QueryActuator(data *model.Actuator, par param.ActuatorQueryParam) int64 {
	count := dao.QueryActuatorSingle(data, par.DeviceTag, par.Tag)

	return count
}

// DeleteActuator
// @Description: 删除执行器
// @param par 查询参数
// @return error
func DeleteActuator(par param.ActuatorQueryParam) error {
	err, counts := dao.DeleteActuator(par.DeviceTag, par.Tag)

	// 未找到执行器
	if counts == 0 {
		return gin.Error{}
	}

	return err
}
