package param

// SensorCreateParam
// @Description: 创建需要的参数
type SensorCreateParam struct {
	Tag       string `json:"tag"`
	Name      string `json:"name"`
	DeviceTag string `json:"device_tag"`
	Unit      string `json:"unit"`
}

// SensorUpdateParam
// @Description: 更新需要的参数
type SensorUpdateParam struct {
	Tag       string `json:"tag"`
	Value     string `json:"value"`
	DeviceTag string `json:"device_tag"`
}

// SensorQueryParam
// @Description: 查询和删除需要的参数
type SensorQueryParam struct {
	DeviceTag string `json:"device_tag"`
	Tag       string `json:"tag"`
}
