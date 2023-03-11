package service

import (
	"github.com/gin-gonic/gin"
	"iot_backend/dao"
	"iot_backend/model"
	"iot_backend/param"
	"iot_backend/tcp"
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

	// 找到执行器，将其历史命令数据全部删除
	dao.DeleteHistoryOrder(par.Tag)

	return err
}

// OrderActuator
// @Description: 执行器下发命令
// @param par
// @return error
func OrderActuator(par param.ActuatorOrderParam) error {
	// 检查是否存在该执行器
	counts := dao.QueryActuatorSingle(&model.Actuator{}, par.DeviceTag, par.Tag)
	if counts == 0 {
		return gin.Error{}
	}

	// 下发命令
	order := "{\"tag\":\"" + par.Tag + "\",\"data\":\"" + par.Value + "\"}"
	err := tcp.SendOrder(order, par.DeviceTag) // 转发给 TCP 业务逻辑去发出命令

	if err == nil {
		// 在历史命令数据里面新增一条记录
		historyOrder := model.HistoryOrderData{
			DeviceTag:   par.DeviceTag,
			ActuatorTag: par.Tag,
			Value:       par.Value,
		}
		CreateHistoryOrder(historyOrder)
	}

	return err
}
