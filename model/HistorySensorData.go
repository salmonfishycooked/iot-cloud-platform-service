package model

// HistorySensorData 历史传感数据表 映射模型
type HistorySensorData struct {
	ID        uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	DeviceTag string `gorm:"type:VARCHAR(255);not NULL" json:"device_tag"`
	SensorTag string `gorm:"type:VARCHAR(255);not NULL" json:"sensor_tag"`
	Value     string `gorm:"type:VARCHAR(255);" json:"value"`
	Created   int64  `gorm:"autoCreateTime" json:"created"`
}
