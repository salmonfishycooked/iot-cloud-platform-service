package tcp

// DeviceTag
// @Description: 设备鉴权结构体
type DeviceTag struct {
	Tag string `json:"tag"`
}

// SensorData
// @Description: 传感数据结构体
type SensorData struct {
	SensorTag string `json:"sensor_tag"`
	Value     string `json:"value"`
}
