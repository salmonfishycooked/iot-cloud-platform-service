package param

import "iot_backend/model"

// DeviceParam
// @Description: 前端发来的参数结构体模型
type DeviceParam struct {
	Tag string `json:"tag"`
}

// DeviceCreateParam
// @Description: 创建需要的参数
type DeviceCreateParam struct {
	Tag  string `json:"tag"`
	Name string `json:"name"`
}

// DeviceResponse
// @Description: 后端发送给前端的结构体模型
type DeviceResponse struct {
	Device    model.Device     `json:"device"`
	Sensors   []model.Sensor   `json:"sensors"`
	Actuators []model.Actuator `json:"actuators"`
}
