package param

// HistorySensorParam
// @Description: 获取历史传感数据参数结构体
type HistorySensorParam struct {
	DeviceTag string `json:"device_tag"`
	SensorTag string `json:"sensor_tag"`
}

// HistoryActuatorParam
// @Description: 获取历史命令数据参数结构体
type HistoryActuatorParam struct {
	DeviceTag   string `json:"device_tag"`
	ActuatorTag string `json:"actuator_tag"`
}
