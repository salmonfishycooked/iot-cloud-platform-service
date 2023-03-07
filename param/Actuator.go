package param

// ActuatorCreateParam
// @Description: 创建需要的参数
type ActuatorCreateParam struct {
	Tag       string `json:"tag"`
	Name      string `json:"name"`
	DeviceTag string `json:"device_tag"`
}

// ActuatorOrderParam
// @Description: 执行器下发命令需要的参数
type ActuatorOrderParam struct {
	Tag       string `json:"tag"`
	Value     string `json:"value"`
	DeviceTag string `json:"device_tag"`
}

// ActuatorQueryParam
// @Description: 查找执行器需要的参数
type ActuatorQueryParam struct {
	DeviceTag string `json:"device_tag"`
	Tag       string `json:"tag"`
}
