package param

// SensorCreateParam
// @Description: 创建需要的参数
type SensorCreateParam struct {
	Tag       string `json:"tag"`
	Name      string `json:"name"`
	DeviceTag string `json:"device_tag"`
}

// SensorUpdateParam
// @Description: 更新需要的参数
type SensorUpdateParam struct {
	Tag       string `json:"tag"`
	Value     string `json:"value"`
	DeviceTag string `json:"device_tag"`
}
